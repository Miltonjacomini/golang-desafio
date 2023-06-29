package controllers

import (
	"miltonjacomini/golang-desafio/models"

	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.GetTop250Movies())
}
