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
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Asia/Jakarta",
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() != fiber.StatusOK {
				// for in the future need to call back using slack or webhook
			}
		},
	}))
	app.Post("/reconciliation", handlerReconciliation.Reconciliation)
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", cfg.Server.HTTP.Port)))
	return nil
}
