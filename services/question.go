package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Question struct {
	Title string
	Body  []byte
}

func (q *Question) save() error {
	filename := q.Title + ".txt"
	return ioutil.WriteFile(filename, q.Body, 0600)
}

func loadPage(title string) (*Question, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Question{Title: title, Body: body}, nil
}

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " rqst found from %s!", r.URL.Path[1:])
}