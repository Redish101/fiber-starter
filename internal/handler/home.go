package handler

import (
	"github.com/Redish101/fiber-starter/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return utils.Res(c, true, "Fiber Starter成功启动", nil)
}