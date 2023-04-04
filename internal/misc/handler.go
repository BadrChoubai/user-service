package misc

import "github.com/gofiber/fiber/v2"

type MiscHandler struct{}

func NewMiscHandler(router fiber.Router) {
	handler := &MiscHandler{}

	router.Get("/health", handler.Healthy)
}

func (handler *MiscHandler) Healthy(ctx *fiber.Ctx) error {
	return ctx.SendString("Healthy")

}
