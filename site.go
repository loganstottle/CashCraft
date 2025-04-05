package main

import (
//    "log"
    "fmt"
    "github.com/gofiber/fiber/v2"
)

func main() {
	config := fiber.Config{
		ServerHeader: "Nmagic's Server",
	}
    app := fiber.New(config)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello oworld!")
    })

	app.Static("/main", "./views/home")
    app.Get("/:skibidi", func(c *fiber.Ctx) error {
        return c.SendString(c.Params("skibidi"))
    })


	name := "pizza"
	fmt.Printf("Hi %s", name)
    fmt.Printf("about to listen")
    app.Listen(":3000")
    fmt.Printf("listening")
}

