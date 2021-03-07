package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

const port = ":8000"

func main() {
	fmt.Printf("Listening on localhost%s\n", port)
	app := fiber.New()
	app.Static("/css", "./public/assets/css")
	app.Static("/js", "./public/assets/js")
	app.Get("/", initVisualiser)
	log.Fatal(app.Listen(port))
}

func initVisualiser(c *fiber.Ctx) error {
	err := c.SendFile("./public/assets/index.html", false)
	return err
}
