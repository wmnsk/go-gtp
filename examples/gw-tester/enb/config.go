// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is a configurations loaded from yaml.
type Config struct {
	LocalAddrs struct {
		S1CIP string `yaml:"s1c_ip"`
		S1UIP string `yaml:"s1u_ip"`
	} `yaml:"local_addresses"`

	MMEAddr  string `yaml:"mme_addr"`
	PromAddr string `yaml:"prom_addr"`

	MCC          string        `yaml:"mcc"`
	MNC          string        `yaml:"mnc"`
	RATType      uint8         `yaml:"rat_type"`
	TAI          uint16        `yaml:"tai"`
	ECI          uint32        `yaml:"eci"`
	Subscribers  []*Subscriber `yaml:"subscribers"`
	UseKernelGTP bool          `yaml:"use_kernel_gtp"`
}

// Subscriber represents a subscriber.
type Subscriber struct {
	IMSI   string `yaml:"imsi"`
	MSISDN string `yaml:"msisdn"`
	IMEISV string `yaml:"imeisv"`
	SrcIP  string `yaml:"src_ip"`
	ITEI   uint32 `yaml:"i_tei"`

	TrafficType string `yaml:"type"`
	EUuIFName   string `yaml:"euu_if_name"`
	HTTPURL     string `yaml:"http_url"`

	Reattach bool `yaml:"reattach_on_reload"`

	// values for these fields are given from MME.
	sgwAddr string
	otei    uint32
}

// String returns the information of s in string.
func (s *Subscriber) String() string {
	return fmt.Sprintf(
		"IMSI: %s, MSISDN: %s, IMEISV: %s, SrcIP: %s, S-GW: %s, I_TEI: %#08x, O_TEI: %#08x",
		s.IMSI, s.MSISDN, s.IMEISV, s.SrcIP, s.sgwAddr, s.ITEI, s.otei,
	)
}

func loadConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if err := yaml.Unmarshal(buf, c); err != nil {
		return nil, err
	}

	return c, nil
}
