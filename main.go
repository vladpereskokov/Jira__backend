package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"os/signal"
	"syscall"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/db"
)

func main() {
	err := auth.InitKeys()

	if err != nil {
		log.Panic("can not init rsa keys: ", err)
	}

	config, err := configs.FromFile("config.json")

	if err != nil {
		log.Panic("bad configs: ", err)
	}

	connection := db.NewDBConnection(config.Mongo)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		connection.CloseConnection()
		os.Exit(0)
	}()

	mux, err := routes.NewRouter()

	if err != nil {
		log.Panic("can not create router: ", err)
	}

	fmt.Printf("Server started on port %d...\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), mux))
}
