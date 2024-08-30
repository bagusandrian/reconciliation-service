package impl

import (
	"net/http"

	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Reconciliation(c *fiber.Ctx) error {
	req := model.GetDummyRequest{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	data, err := h.usecase.GetDummy(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"data":    data,
			"error":   err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}
