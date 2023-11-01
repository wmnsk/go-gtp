// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command enb works as pseudo eNB that forwards packets through GTPv1 tunnel.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"syscall"
	"fmt"
)

func main() {
	var configPath = flag.String("config", "./enb.yml", "Path to the configuration file.")
	flag.Parse()
	log.SetPrefix("[eNB] ")

	cfg, err := loadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	enb, err := newENB(cfg)
	if err != nil {
		log.Printf("failed to initialize eNB: %s", err)
	}

	IMSI := "001010000000010"
	MSISDN := "0000000010"
	IMEISV := "1234500000010"
	SRCIP := 10
	SRCIP2 := 100
	srcip2 := strconv.Itoa(SRCIP2)
	TrafficType := "http_get"
	EUuIFName := "lo"
	HTTPURL := "http://172.22.0.254/"
	for i := 0; i < 100; i++ {
		
		imsi := imsiGenerator(i, IMSI)
		msisdn := imsiGenerator(i, MSISDN)
		imeisv := imsiGenerator(i, IMEISV)
		srcip := strconv.Itoa(SRCIP)
		if srcip == "255" {
			SRCIP = 1	
			srcip = "1"
			SRCIP2++
			srcip2 = strconv.Itoa(SRCIP2)
		}
		var sub *Subscriber
		sub = &Subscriber{IMSI:imsi, MSISDN:msisdn, IMEISV:imeisv, SrcIP: "192.168."+srcip2+"."+srcip, TrafficType:TrafficType, EUuIFName:EUuIFName, HTTPURL:HTTPURL}
		enb.candidateSubs= append(enb.candidateSubs, sub)
		SRCIP++
	}
	spew.Dump(enb.candidateSubs)

	defer enb.close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fatalCh := make(chan error, 1)
	go func() {
		if err := enb.run(ctx); err != nil {
			fatalCh <- err
		}
	}()

	for {
		select {
		case sig := <-sigCh:
			switch sig {
			case syscall.SIGINT:
				return
			case syscall.SIGHUP:
				// reload config and attach/detach subscribers again
				newCfg, err := loadConfig(*configPath)
				if err != nil {
					log.Printf("Error reloading config %s", err)
				}

				if err := enb.reload(newCfg); err != nil {
					log.Printf("Error applying reloaded config %s", err)
				}
			}
		case err := <-enb.errCh:
			log.Printf("WARN: %s", err)
		case err := <-fatalCh:
			log.Printf("FATAL: %s", err)
			return
		}
	}
}


func imsiGenerator(i int, msin string) string {

	msin_int, err := strconv.Atoi(msin)
	if err != nil {
		log.Fatal("Error in get configuration")
	}
	base := msin_int + (i -1)

	imsi := "00" + fmt.Sprintf("%010d", base)
	return imsi
}