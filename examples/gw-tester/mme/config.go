// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is a configurations loaded from yaml.
type Config struct {
	LocalAddrs struct {
		S1CAddr string `yaml:"s1c_addr"`
		S11IP   string `yaml:"s11_ip"`
	} `yaml:"local_addresses"`

	PromAddr string `yaml:"prom_addr"`

	MCC string `yaml:"mcc"`
	MNC string `yaml:"mnc"`

	APN string `yaml:"apn"`

	SgwS11 string `yaml:"sgw_s11_ip"`
	PgwS5C string `yaml:"pgw_s5c_ip"`
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
