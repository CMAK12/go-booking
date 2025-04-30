package v1

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		log.Printf("Started %s %s", c.Method(), c.Path())

		err := c.Next()

		log.Printf("Completed %s in %v", c.Path(), time.Since(start))
		return err
	}
}

func ResponseWrapper(handler func(c *fiber.Ctx) (any, int, int64, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, status, count, err := handler(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return write(c, status, data, count)
	}
}

func write(ctx *fiber.Ctx, status int, data any, count int64) error {
	response := fiber.Map{
		"data": data,
	}
	if count > 0 {
		response["count"] = count
	}

	return ctx.Status(status).JSON(response)
}
