package main

import (
	"encoding/xml"
	"net/http"
)

// Profile represent user profile to be jsonified
type ProfileXML struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/", indexXML)
	http.ListenAndServe(":3001", nil)
}

func indexXML(w http.ResponseWriter, r *http.Request) {
	profile := ProfileXML{"Andrew", []string{"bicycle", "programming", "books"}}

	js, err := xml.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(js)
}
