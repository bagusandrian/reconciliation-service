package main

import (
	"fmt"
	"log"
	"time"

	hImpl "github.com/bagusandrian/reconciliation-service/internals/handler/http/impl"
	"github.com/bagusandrian/reconciliation-service/internals/model"
	readFileImpl "github.com/bagusandrian/reconciliation-service/internals/repository/readfile/impl"
	ucImpl "github.com/bagusandrian/reconciliation-service/internals/usecase/reconciliation/impl"
	"github.com/gofiber/fiber/v2"
)

func startApp(cfg *model.Config) (err error) {

	log.Println("starting http server", time.Now())
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	repositoryReadFile := readFileImpl.New(cfg)
	usecaseReadFile := ucImpl.New(cfg, repositoryReadFile)
	handlerReconciliation := hImpl.New(cfg, usecaseReadFile)

	app := fiber.New(fiber.Config{
		AppName: "reconciliation Service",
	})
	app.Post("/reconciliation", handlerReconciliation.Reconciliation)
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", cfg.Server.HTTP.Port)))
	return nil
}
