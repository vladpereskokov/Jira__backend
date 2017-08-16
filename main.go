package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/DVI-GI-2017/Jira__backend/auth"
)

func rsaInit() {
	err := auth.InitKeys()

	if err != nil {
		log.Panic("can not init rsa keys: ", err)
	}
}

func startRouter() (mux http.Handler) {
	mux, err := routes.NewRouter()

	if err != nil {
		log.Panic("can not create router: ", err)
	}

	return
}

func main() {
	rsaInit()

	configs.ParseFromFile("config.json")
	mux := startRouter()

	fmt.Printf("Server started on port %d...\n", configs.ConfigInfo.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configs.ConfigInfo.Server.Port), mux))
}
