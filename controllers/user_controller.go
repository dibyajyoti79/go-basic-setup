package controllers

import (
	"pripgo/services"
	"pripgo/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error {
	user, err := services.GetUserService(ctx.Params("id"))
	if err != nil {
		return err
	}
	return ctx.JSON(utils.NewApiResponse(200, user, "User fetched successfully"))
}
