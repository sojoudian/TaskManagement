package main

import (
	"log"
	"net/http"

	"github.com/sojoudian/TaskManagement/task"
)

func main() {
	// register RESTful endpoint handler for '/tasks/'
	http.Handle("/tasks/", &task.TaskAPI{})
	log.Fatalln(http.ListenAndServe(":3003", nil))
}
