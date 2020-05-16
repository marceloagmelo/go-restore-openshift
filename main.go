package main

import (
	"github.com/marceloagmelo/go-restore-openshift/logger"

	"github.com/marceloagmelo/go-restore-openshift/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	logger.Info.Println("Listen 8080...")
	app.Run(":8080")
}
