// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command mme works as pseudo MME/HSS communicates S/P-GW with GTPv2 signaling.
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
	var configPath = flag.String("config", "./mme.yml", "Path to the configuration file.")
	flag.Parse()
	log.SetPrefix("[MME] ")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP)

	cfg, err := loadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	mme, err := newMME(cfg)
	if err != nil {
		log.Printf("failed to initialize MME: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fatalCh := make(chan error)
	go func() {
		if err := mme.run(ctx); err != nil {
			fatalCh <- err
		}
	}()

	for {
		select {
		case sig := <-sigCh:
			switch sig {
			case syscall.SIGINT:
				cancel()
				return
			case syscall.SIGHUP:
				// reload config and attach/detach subscribers
				newCfg, err := loadConfig(*configPath)
				if err != nil {
					log.Printf("Error reloading config %s", err)
				}

				if err := mme.reload(newCfg); err != nil {
					log.Printf("Error applying reloaded config %s", err)
				}
			}
		case err := <-mme.errCh:
			log.Printf("WARN: %s", err)
		case err := <-fatalCh:
			log.Printf("FATAL: %s", err)
			return
		}
	}
}
