package task

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type TaskAPI struct {
}
type Task struct {
	ID        uint64 `json:"id,omitempty"`
	TaskTitle string `json:"task_title,omitempty"`
	TaskBody  string `json:"task_body,omitempty"`
}

var db = []*Task{}

//mutex- a mutual exclusion
var nextTaskID uint64
var lock sync.Mutex

func (t *TaskAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	case http.MethodPost:
		doPost(w, r)
	case http.MethodDelete:
		TDelete(w, r)

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to '%v' \n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to '%v' \n", r.Method, r.URL)

	}
}
