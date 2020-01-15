// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command pgw is a dead simple implementation of P-GW only with GTP-related features.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configPath = flag.String("config", "./pgw.yml", "Path to the configuration file.")
	flag.Parse()
	log.SetPrefix("[P-GW] ")

	cfg, err := loadConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}

	pgw, err := newPGW(cfg)
	if err != nil {
		log.Printf("failed to initialize P-GW: %s", err)
		return
	}
	defer pgw.close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fatalCh := make(chan error)
	go func() {
		if err := pgw.run(ctx); err != nil {
			fatalCh <- err
		}
	}()

	for {
		select {
		case sig := <-sigCh:
			// TODO: reload config on receiving SIGHUP
			log.Println(sig)
			return
		case err := <-pgw.errCh:
			log.Printf("WARN: %s", err)
		case err := <-fatalCh:
			log.Printf("FATAL: %s", err)
			return
		}
	}
}
