package utils

import "github.com/gofiber/fiber/v2"

type res struct {
	Ok bool `json:"ok"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Res(c *fiber.Ctx, ok bool, msg string, data interface{}) error {
	return c.JSON(res{
		Ok: ok,
		Msg: msg,
		Data: data,
	})
}