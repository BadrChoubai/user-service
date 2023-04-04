package service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(baseServiceRoute fiber.Router, userService UserService) {
	handler := &UserHandler{
		userService: userService,
	}

	baseServiceRoute.Use(limiter.New(limiter.Config{
		Max:        3,
		Expiration: 10 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))

	baseServiceRoute.Get("/:userId", handler.GetSingleUser)
}

func (handler *UserHandler) GetSingleUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userId, err := c.ParamsInt("userId")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user id",
		})
	}

	user, err := handler.userService.GetUserById(customContext, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})

}
