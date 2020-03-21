package config

import (
	"html/template"
	"net/http"
)

// Templates obj for app.
var Templates = template.Must(template.ParseGlob("templates/*.html"))

// BlogURL Url for the app
const BlogURL string = "http://go_todo:80"

// MongodbURL for the mongo API
const MongodbURL string = "http://mongo_todo:3000/todo"

// Client http.client pointer
var Client = &http.Client{}
