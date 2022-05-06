package main

import (
	"github.com/sojoudian/TaskManagement/cmd/API/task"
	"log"
	"net/http"
)

func main() {
	// register RESTful endpoint handler for '/tasks/'
	http.HandleFunc("/tasks", &task.TaskAPI{})
	log.Fatalln(http.ListenAndServe("3003", nil))
}
