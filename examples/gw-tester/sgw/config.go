// Copyright 2019-2023 go-gtp authors. All rights reserved.
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
		S11IP string `yaml:"s11_ip"`
		S1UIP string `yaml:"s1u_ip"`
		S5CIP string `yaml:"s5c_ip"`
		S5UIP string `yaml:"s5u_ip"`
	} `yaml:"local_addresses"`

	UseKernelGTP bool   `yaml:"use_kernel_gtp"`
	PromAddr     string `yaml:"prom_addr"`
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
