package http

import (
	"github.com/gofiber/fiber/v2"
)

//go:generate mockery --name=Handler --filename=mock_handler.go --inpackage
type Handler interface {
	Reconciliation(c *fiber.Ctx) error
}
