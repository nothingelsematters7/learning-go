package main

import (
	"encoding/json"
	"net/http"
)

// Profile represent user profile to be jsonified
type ProfileJSON struct {
	Name    string
	Hobbies []string
}

// func main() {
// 	http.HandleFunc("/", indexJSON)
// 	http.ListenAndServe(":3000", nil)
// }

func indexJSON(w http.ResponseWriter, r *http.Request) {
	profile := ProfileJSON{"Andrew", []string{"bicycle", "programming", "books"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
