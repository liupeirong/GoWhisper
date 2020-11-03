package main

import (
  "encoding/json"
  "net/http"
	"log"
	"bytes"
	"os"
	"github.com/gorilla/mux"
)

type Gossip struct {
	Sender string `json:"sender"`
	Message string `json:"message"`
}

var gossips []Gossip

func GetGossipsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(gossips)
}

func ReceiveGossipHandler(w http.ResponseWriter, r *http.Request) {
	var gossip Gossip
	json.NewDecoder(r.Body).Decode(&gossip)
	gossips = append(gossips, gossip)
	if gossip.Sender == os.Getenv("ENV_MYSELF") {
		StopGossip(w, gossip)
	} else { 
		PassGossip(w, gossip)
	}
}

func StopGossip(w http.ResponseWriter, gossip Gossip) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("that's myself"))
}

func PassGossip(w http.ResponseWriter, gossip Gossip) {
	jsonReq, err := json.Marshal(gossip)
	resp, err := http.Post(os.Getenv("ENV_FORWARDURL"), "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
  if err != nil {
      log.Fatalln(err)
  }
	defer resp.Body.Close()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gossips", GetGossipsHandler).Methods("GET")
	r.HandleFunc("/whisper", ReceiveGossipHandler).Methods("POST")
	http.ListenAndServe(":5000", r)
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}