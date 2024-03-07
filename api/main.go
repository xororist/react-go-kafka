package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MessageResponse struct {
	uuid    uuid.UUID
	Content string `json:"content"`
	date    time.Time
}

type MessageRequest struct {
	Content string `json:"content"`
	Date    string `json:"date"`
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber-kafka",
		AppName:       "fiber-fafka-api v0.0.1",
	})

	var messages []MessageResponse

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(messages)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		message := new(MessageRequest)

		if err := c.BodyParser(message); err != nil {
			return err
		}

		data := MessageResponse{
			uuid:    uuid.New(),
			Content: message.Content,
			date:    time.Now(),
		}

		messages = append(messages, data)

		return c.SendStatus(fiber.StatusCreated)
	})

	app.Listen(":4000")
}
