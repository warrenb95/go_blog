package controllers

import (
	"net/http"

	"github.com/warrenb95/go_blog/client/config"
)

// IndexEndPoint The Index page end point to serve up a list of blog posts newest -> oldest
func IndexEndPoint(res http.ResponseWriter, req *http.Request) {
	err := config.Templates.ExecuteTemplate(res, "index", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
