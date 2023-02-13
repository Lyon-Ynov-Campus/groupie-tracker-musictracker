package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	Star := template.Must(template.ParseGlob("./static/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Music Tracker")
		Star.ExecuteTemplate(w, "index.html", "aa")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}