package main

import (
	c "golang-socket/client"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)


func main() {
	inputs := []string{}
	outputs := []string{}
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})


	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello, World! sss",
		})
	})

	app.Get("/go-client/:name", func(ctx *fiber.Ctx) error {
		input := ctx.Params("name")
		output := c.Client(input)

		inputs = append(inputs, input)
		outputs = append(outputs, output)

		return ctx.Render("index", fiber.Map{
			"Inputs": inputs,
			"Outputs": outputs,
		})
	})

	app.Listen(":3000")


}
