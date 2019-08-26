package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hi everyone!")
}

func main() {
	router := httprouter.New()

	router.GET("/sayhello", sayHello)

	fmt.Println("Listening @ localhost:8080....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
