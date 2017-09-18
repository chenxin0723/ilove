package utils

import (
	"html/template"
	"net/http"
)

func FuncMap(w http.ResponseWriter, req *http.Request) template.FuncMap {
	funcMap := template.FuncMap{}
	return funcMap
}
