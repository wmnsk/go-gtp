// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command sgw is a dead simple implementation of S-GW.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func main() {
	var configPath = flag.String("config", "./sgw.yml", "Path to the configuration file.")
	flag.Parse()
	log.SetPrefix("[S-GW] ")

	cfg, err := loadConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}

	sgw, err := newSGW(cfg)
	if err != nil {
		log.Printf("failed to initialize SGW: %s", err)
	}
	defer sgw.close()

	// register handlers for ALL the messages you expect remote endpoint to send.
	sgw.s11Conn.AddHandlers(map[uint8]v2.HandlerFunc{
		messages.MsgTypeCreateSessionRequest: sgw.handleCreateSessionRequest,
		messages.MsgTypeModifyBearerRequest:  sgw.handleModifyBearerRequest,
		messages.MsgTypeDeleteSessionRequest: sgw.handleDeleteSessionRequest,
		messages.MsgTypeDeleteBearerResponse: sgw.handleDeleteBearerResponse,
	})
	sgw.s5cConn.AddHandlers(map[uint8]v2.HandlerFunc{
		messages.MsgTypeCreateSessionResponse: sgw.handleCreateSessionResponse,
		messages.MsgTypeDeleteSessionResponse: sgw.handleDeleteSessionResponse,
		messages.MsgTypeDeleteBearerRequest:   sgw.handleDeleteBearerRequest,
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fatalCh := make(chan error)
	go func() {
		if err := sgw.run(ctx); err != nil {
			fatalCh <- err
		}
	}()

	for {
		select {
		case sig := <-sigCh:
			// TODO: reload config on receiving SIGHUP
			log.Println(sig)
			return
		case err := <-sgw.errCh:
			log.Printf("WARN: %s", err)
		case err := <-fatalCh:
			log.Printf("FATAL: %s", err)
			return
		}
	}
}
