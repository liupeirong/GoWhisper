package main

import (
  "encoding/json"
  "net/http"
)

type Gossip struct {
	Sender string `json:"Sender"`
	Message string `json:"Message"`
}

func GetGossipsHandler(w http.ResponseWriter, r *http.Request) {
 
	gossips := []Gossip{
			Gossip{"John", "This is first post."},
			Gossip{"Jane", "This is second post."},
	}

	json.NewEncoder(w).Encode(gossips)
}

func main() {
	http.HandleFunc("/gossips", GetGossipsHandler)
	http.ListenAndServe(":5000", nil)
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}