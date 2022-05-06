package task

import (
	"fmt"
	"net/http"
)

type TaskAPI struct {
	Todo string
}

func (t *TaskAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Go call from %v, ", r.URL)
}
