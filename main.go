package main

import (
	"dowwnload-file/src/controller"
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must((template.ParseGlob("src/views/*.html")))
}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index", nil)

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/download", controller.Index)
	http.ListenAndServe(":"+port, mux)

}
