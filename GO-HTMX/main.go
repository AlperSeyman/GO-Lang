package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct{
	Title string
	Director string
}

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			
		}

		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
