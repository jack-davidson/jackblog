package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (b Blog) FindByTitle(title string) (Article, error) {
	for _, article := range b {
		if article.Title == title {
			return article, nil
		}
	}
	return Article{}, errors.New(fmt.Sprintf("could not find article with .Title='%s'", title))
}

func (b Blog) FindByRouteTitle(routeTitle string) (Article, error) {
	for _, article := range b {
		if article.RouteTitle == routeTitle {
			return article, nil
		}
	}
	return Article{}, errors.New(fmt.Sprintf(
		"could not find article with .RouteTitle='%s'", routeTitle))
}
