package main

import (
	"net/http"

	"github.com/angadthandi/goapp/dbconnect"
	"github.com/angadthandi/goapp/gosocket"
	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/registry"
	"github.com/angadthandi/goapp/route"
	"github.com/angadthandi/goapp/util/config"
)

func main() {
	// read config
	configData, err := config.Read()
	if err != nil {
		log.Fatal("unable to read config: %v", err)
	}

	// set logger
	log.SetLogLevel(configData.LogLevel)
	log.SetOutputFile(
		configData.LogPath,
		configData.LogMaxSize,
		configData.LogMaxBackups,
		configData.LogMaxAge,
		configData.LogCompress,
	)

	// connect database
	dbRef := dbconnect.Conn(configData)
	log.Info("Initialized database!")
	defer dbRef.Close()

	// initialize registry
	clientRegistry := registry.NewRegistry()

	// start hub
	// for creating websocket conns
	hub := gosocket.NewHub()
	go hub.Run(clientRegistry)

	// initialize routes handler
	route.Handle(dbRef, hub, clientRegistry, configData)

	log.Infof("Listening on Port: %v", configData.ServerPort)
	// start http web server
	log.Fatal(http.ListenAndServe(":"+configData.ServerPort, nil))
}
