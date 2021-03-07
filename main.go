package main

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber"
)

const port = ":8000"

func main() {
	fmt.Printf("Listening on localhost%s\n", port)
	app := fiber.New()
	app.Static("/static", "./public/assets/css")
	app.Get("/", initVisualiser)
	app.Listen(port)
}

func initVisualiser(c *fiber.Ctx) error {
	err := c.SendFile("./public/assets/index.html", false)
	return err
}

func generateArray(maxValue int, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result = append(result, rand.Intn(length))
	}
	return result
}

func populateVisualisation(height []int) {
	// for index, height := range height {

	// }
}
