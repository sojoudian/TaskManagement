package task

import (
	"fmt"
	"log"
	"net/http"
)

type TaskAPI struct {
}
type Task struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var db []*Task

func (t *TaskAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to '%v' \n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to '%v' \n", r.Method, r.URL)

	}
}
