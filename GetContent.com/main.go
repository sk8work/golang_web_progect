package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Println("listening on port :4000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.Handle("/jpg/", http.StripPrefix("/jpg/", http.FileServer(http.Dir("./jpg/"))))

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":4000", nil)
}
