package task

import (
	"encoding/json"
	"net/http"
)

func doGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(db)
	//fmt.Printf("Go call from %v, using HTTP METHOD %v \n", r.URL, r.Method)
}
