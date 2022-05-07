package task

import (
	"encoding/json"
	"net/http"
)

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonDecoder := json.NewDecoder(r.Body)

	aTask := &Task{}
	err := jsonDecoder.Decode(aTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start of protected code changes
	lock.Lock()
	nextTaskID++
	aTask.ID = nextTaskID
	db = append(db, aTask)
	lock.Unlock()
	// end of protected code changes

	respTask := Task{
		ID:        aTask.ID,
		TaskTitle: aTask.TaskTitle,
		TaskBody:  aTask.TaskBody,
	}
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(respTask)
	//fmt.Printf("Go call from %v, using HTTP METHOD %v \n", r.URL, r.Method)
}
