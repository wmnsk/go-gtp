// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metricsCollector struct {
	activeSessions   prometheus.GaugeFunc
	messagesSent     *prometheus.CounterVec
	messagesReceived *prometheus.CounterVec
}

func (m *mme) runMetricsCollector() error {
	mc := &metricsCollector{}
	mc.activeSessions = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mme_active_sessions",
			Help: "number of session established currently",
		},
		func() float64 {
			return float64(m.s11Conn.SessionCount())
		},
	)

	mc.messagesSent = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mme_messages_sent_total",
			Help: "number of messages sent by messagge type",
		},
		[]string{"dst", "type"},
	)

	mc.messagesReceived = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mme_messages_received_total",
			Help: "number of messages received by messagge type",
		},
		[]string{"src", "type"},
	)

	m.mc = mc
	return nil
}
