package main

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/gomarkdown/markdown"
	_ "github.com/lib/pq"
)

type Article struct {
	Title      string        `json:"title"`
	RouteTitle string        `json:"routetitle"`
	Author     string        `json:"author"`
	Body       template.HTML `json:"body"`
	Date       string        `json:"date"` /* (month)/(day)/(year) */
}

type Blog []Article

func NewBlog() Blog {
	blog := Blog{}
	data, _ := os.ReadFile("blog.json")
	json.Unmarshal(data, &blog)
	for _, article := range blog {
		article.Body = template.HTML(markdown.ToHTML([]byte(article.Body), nil, nil)[:])
	}
	return blog
}
