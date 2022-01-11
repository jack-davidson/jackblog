package main

import (
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		CacheDuration: 1 * time.Microsecond,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		c.Context().SetContentType("text/html")
		tmpl.Execute(c, nil)

		return nil
	})

	app.Listen(":3000")
}
