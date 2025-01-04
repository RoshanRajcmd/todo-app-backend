package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Name string
}

func main() {

	var router = http.NewServeMux()

	//make Http routes. If you dont enter the specific type of request then the
	//complier will consider as all type of request
	router.HandleFunc("/hi", func(writter http.ResponseWriter, resp *http.Request) {
		//code implementation for the route
		//Fprint and writter obj is used to write in the route page
		fmt.Fprintf(writter, "Hello World")
		//The Below line can also be used for same
		writter.Write([]byte("Hello World"))
	})

	router.HandleFunc("GET /hi/{name}", func(w http.ResponseWriter, r *http.Request) {
		var name = r.PathValue("name")
		var tmpl = template.Must(template.New("").ParseGlob("./index/*"))
		tmpl.ExecuteTemplate(w, "index.html", PageData{
			Name: name,
		})
		//fmt.Fprintf(w, "Hello World", name)
	})

	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		var name = r.URL.Query().Get("name")
		w.Write([]byte("Hello: " + name))
	})

	var server = http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Server is runner on port" + server.Addr)

	//Start and serve the server in the given port
	//If there is any issue there below function will return a error
	log.Fatal(server.ListenAndServe())
}
