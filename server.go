package main

import (
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	TemplateRoot = "templates"
)

func Render(c *fiber.Ctx, name string, data interface{}) error {
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
		Render(c, "index", NewBlog())
		return nil
	})

	app.Get("/blogs/:title", func(c *fiber.Ctx) error {
		fmt.Println(c)
		b := NewBlog()
		article, err := b.FindByRouteTitle(c.Params("title"))
		if err != nil {
			return errors.New("404 Not Found")
		}
		Render(c, "blog", article)
		return nil
	})

	app.Get("/blogs/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		c.Redirect("/")
		return nil
	})

	app.Get("/new/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		Render(c, "newblog", nil)
		return nil
	})

	app.Post("/new/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		return nil
	})

	app.Listen(":3000")
}
