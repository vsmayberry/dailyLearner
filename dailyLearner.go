package main

import (
	"./services"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/question/", services.QuestionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
