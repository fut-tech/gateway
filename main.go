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

import (
	"fmt"

	"github.com/fut-tech/gateway/utils"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(utils.NewLog)
	container.Provide(utils.NewConfigs)

	container.Provide(NewApp)

	err := container.Invoke(func(app *App) {
		app.Start()
	})

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
	}
}
