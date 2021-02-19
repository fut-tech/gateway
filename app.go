////////////////////////////////////////////////////////////////////
//
// Future Technologies Inc.
//
// Copyright (C) 2021 Sergey Denisov. GPLv3
//
// Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
//
////////////////////////////////////////////////////////////////////

package main

import "github.com/fut-tech/gateway/utils"

// App Application struct
type App struct {
	log *utils.Log
	cfg *utils.Configs
}

// NewApp Make new struct
func NewApp(l *utils.Log, c *utils.Configs) *App {
	return &App{
		log: l,
		cfg: c,
	}
}

// Start Start application
func (a *App) Start() {
	// Set log path
	a.log.SetPath("/var/log/ft/")

	// Loading app configs
	var err = a.cfg.LoadFromFile("/etc/ft/gateway.conf")
	if err != nil {
		err = a.cfg.LoadFromFile("gateway.conf")
		if err != nil {
			a.log.Error("APP", "Fail to load configs", err)
			return
		}
	}
}
