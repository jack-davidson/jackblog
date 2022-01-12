package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gomarkdown/markdown"
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

func LoadBlogs(name string) []Article {
	articles := []Article{}
	data, _ := os.ReadFile(name)
	json.Unmarshal(data, &articles)
	return articles
}

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		CacheDuration: 1 * time.Microsecond,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		RenderTemplate(c, "index", LoadBlogs("blog.json"))
		return nil
	})

	app.Get("/blogs/:title", func(c *fiber.Ctx) error {
		for _, article := range LoadBlogs("blog.json") {
			if article.RouteTitle == c.Params("title") {
				article.Body = template.HTML(markdown.ToHTML([]byte(article.Body), nil, nil)[:])
				RenderTemplate(c, "blog", article)
			}
		}
		return nil
	})

	app.Get("/blogs/", func(c *fiber.Ctx) error {
		c.Redirect("/")
		return nil
	})

	app.Listen(":3000")
}
