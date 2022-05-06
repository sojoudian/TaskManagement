package task

import (
	"fmt"
	"net/http"
)

func doGet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Go call from %v, using HTTP METHOD %v \n", r.URL, r.Method)
}
