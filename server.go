package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	TemplateRoot = "templates"
)

func RenderTemplate(c *fiber.Ctx, name string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/%s.html", TemplateRoot, name))
	if err != nil {
		return err
	}
	c.Context().SetContentType("text/html")
	return tmpl.Execute(c, data)
}

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		CacheDuration: 1 * time.Microsecond,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		RenderTemplate(c, "index", nil)
		return nil
	})

	app.Get("/blogs/:title", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("title"))
		RenderTemplate(c, "blog", nil)
		return nil
	})

	app.Get("/blogs/", func(c *fiber.Ctx) error {
		c.Redirect("/")
		return nil
	})

	app.Listen(":3000")
}
