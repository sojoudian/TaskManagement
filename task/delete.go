package task

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TDelete - deletes a task from the db using the path '/tasks/id', eg: /tasks/2
func TDelete(w http.ResponseWriter, r *http.Request) {
	// get the user ID from the path
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to delete user %v", id)

	// start of protected code changes
	lock.Lock()
	var tmp = []*Task{}
	for _, t := range db {
		if id == t.ID {
			continue
		}
		tmp = append(tmp, t)
	}
	db = tmp
	// end of protected code changes
	lock.Unlock()

	fmt.Println(fields)
	w.Header().Set("Content-Type", "application/json")
}
