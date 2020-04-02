package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Human struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/index.html"))
	tmpl.Execute(w, "tokutatsu")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	my := Human{
		Name: "tokutatsu",
		Age:  21,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(my)
}
