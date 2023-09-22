package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Zmey56 struct {
	Name   string
	Skills string
}

func getZmey56(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(zmey56)
}

var zmey56 Zmey56

func createZmey56(ctx *fiber.Ctx) error {
	body := new(Zmey56)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	zmey56 = Zmey56{
		Name:   body.Name,
		Skills: body.Skills,
	}

	return ctx.Status(fiber.StatusOK).JSON(zmey56)
}

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	zmey56App := app.Group("/zmey56")
	zmey56App.Get("", getZmey56)
	zmey56App.Post("", createZmey56)

	app.Listen(":80")
}
