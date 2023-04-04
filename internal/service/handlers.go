package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type UserHandler struct {
	userService UserService
}

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
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

// @Description	Get a single user
// @Id				get-user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			userId	path		int	true	"User Id"
// @Success		200		{object}	ResponseHTTP{data=User}
// @Failure		404		{object}	ResponseHTTP{}
// @Router			/users/{userId} [get]
func (handler *UserHandler) GetSingleUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userId, err := c.ParamsInt("userId")
	fmt.Println(userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Data:    "Please specify a valid user id",
		})
	}

	user, err := handler.userService.GetUserById(customContext, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Data:    user,
	})

}
