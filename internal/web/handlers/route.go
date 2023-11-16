package handlers

import "github.com/gofiber/fiber/v2"

type Route interface {
	Pattern()
	Handle(ctx *fiber.Ctx)
}
