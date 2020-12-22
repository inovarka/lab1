package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type indexResult struct {
	Time string `json:"time"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := indexResult{time.Now().Format(time.RFC3339)}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonRes)
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8795", nil)
}
