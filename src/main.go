package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"localhost/go-todo/src/controllers"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", controllers.TodoIndex)
	router.GET("/todos/:todoId", controllers.TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello %q\n", html.EscapeString(r.URL.Path))
}
