package misc

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type MiscHandler struct{}

func NewMiscHandler(router fiber.Router) {
	handler := &MiscHandler{}

	router.Get("/health", handler.Healthy)
	router.Get("/swagger/*", swagger.HandlerDefault)
}

func (handler *MiscHandler) Healthy(ctx *fiber.Ctx) error {
	return ctx.SendString("Healthy")
}
