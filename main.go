package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func handler(c *fiber.Ctx) error {
	fmt.Fprintf(c, "Arroba")
	return nil
}

type account struct {
	User string 
	Password string 
	Email string 
}

func main() {
	app := fiber.New()
	app.Get("/", handler)
	app.Listen(":8080")
}