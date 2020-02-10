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
		S5CIP string `yaml:"s5c_ip"`
		S5UIP string `yaml:"s5u_ip"`
		SGiIP string `yaml:"sgi_ip"`
	} `yaml:"local_addresses"`

	PromAddr string `yaml:"prom_addr"`

	SGiIFName   string `yaml:"sgi_if_name"`
	RouteSubnet string `yaml:"route_subnet"`
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
