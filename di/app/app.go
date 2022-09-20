package main

import (
	"di"
	"fmt"
	"log"
	"net/http"
	"os"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Page accessed.")
	di.Greet(w, "WRLD")
}

func main() {
	fmt.Println("Hello world!")
	di.Greet(os.Stdout, "Mike")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
