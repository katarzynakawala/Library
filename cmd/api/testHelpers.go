package main

import (
	"log"
	"os"
)

func applicationInstance(port int, env string) *application {
	var cfg config
	cfg.port = port
	cfg.env = env

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}
	return app
}