////////////////////////////////////////////////////////////////////
//
// Future Technologies Inc.
//
// Copyright (C) 2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
////////////////////////////////////////////////////////////////////

package utils

import (
	"encoding/json"
	"io/ioutil"
)

type ServerCfg struct {
	IP   string
	Port int
}

type AppCfg struct {
	Server ServerCfg
}

// Configs App configs
type Configs struct {
	configs AppCfg
}

// NewConfigs make new struct
func NewConfigs() *Configs {
	return &Configs{}
}

// LoadFromFile Loading configs from file
func (c *Configs) LoadFromFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c.configs)
	if err != nil {
		return err
	}

	return nil
}

// Settings Get app configs
func (c *Configs) Settings() *AppCfg {
	return &c.configs
}
