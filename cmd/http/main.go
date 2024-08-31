package main

import (
	"log"

	"github.com/bagusandrian/reconciliation-service/internals/config"
)

var cfgTest bool

const (
	repoName = "reconciliation-service"
	appName  = repoName + "-http"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("%v", err)
		}
	}()
	log.Println("Starting reconciliation-service http")

	// init config
	cfg, err := config.New(repoName)
	if err != nil {
		log.Fatalf("failed to init the config: %v", err)
	}
	log.Println("init config done")

	err = startApp(cfg)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
