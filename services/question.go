package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	if strings.Compare(r.Method, "GET") == 0 {
		fmt.Fprintf(w, "GET rqst found from %s!", r.URL.Path[1:])
	} else if strings.Compare(r.Method, "POST") == 0 {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "POST rqst found from %s! unable to parse Form", r.URL.Path[1:])
		}
		fmt.Fprintf(w, "POST rqst found from %s with body! %s", r.URL.Path[1:], r.PostForm)
	}
	
}