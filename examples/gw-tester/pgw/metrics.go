// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/vishvananda/netlink"
)

type metricsCollector struct {
	activeSessions   prometheus.GaugeFunc
	activeBearers    prometheus.GaugeFunc
	messagesSent     *prometheus.CounterVec
	messagesReceived *prometheus.CounterVec
}

func (p *pgw) runMetricsCollector() error {
	mc := &metricsCollector{}
	mc.activeSessions = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pgw_active_sessions",
			Help: "number of session established currently",
		},
		func() float64 {
			return float64(p.cConn.SessionCount())
		},
	)

	mc.activeBearers = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pgw_active_bearers",
			Help: "number of GTP-U tunnels established currently",
		},
		func() float64 {
			tunnels, err := netlink.GTPPDPList()
			if err != nil {
				log.Printf("metrics: could not get tunnels: %s", err)
				return 0
			}
			return float64(len(tunnels))
		},
	)

	mc.messagesSent = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pgw_messages_sent_total",
			Help: "number of message sent by messagge type",
		},
		[]string{"dst", "type"},
	)

	mc.messagesReceived = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pgw_messages_received_total",
			Help: "number of message received by messagge type",
		},
		[]string{"src", "type"},
	)

	p.mc = mc
	return nil
}
