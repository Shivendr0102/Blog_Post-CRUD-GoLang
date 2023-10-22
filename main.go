package main

import (
	"22nd_Oct_Antino/config"
	"22nd_Oct_Antino/db"
	"22nd_Oct_Antino/routes"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println(("Get set Go"))

	configuration, err := config.New()
	if err != nil {
		log.Error("Configuration error: %s", err)
	}

	dbSource, err := db.NewSource(configuration)
	if err != nil {
		log.Error("Failed to connect to db source")
	}

	router := routes.New(configuration, dbSource)

	if configuration.Constants.AllowHttp {
		go serveAPIinHttp(configuration, router)
	}
}

func serveAPIinHttp(configuration *config.Config, router *chi.Mux) {
	log.Error("Serving application at :%s", configuration.Constants.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.Port, router))
}
