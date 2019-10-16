package main

import (
	"fmt"
	"github.com/vsmayberry/dailyLearner/services"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listening on ::8080")
	http.HandleFunc("/question/", services.QuestionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
