package main

import "html/template"

type Article struct {
	Title      string        `json:"title"`
	RouteTitle string        `json:"routetitle"`
	Author     string        `json:"author"`
	Body       template.HTML `json:"body"`
	Date       string        `json:"date"` /* (month)/(day)/(year) */
}
