package main

import (
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
)

func TestReceiveGossipHandler_WhenMyself_StopGossip(t *testing.T) {
	//setup
	os.Setenv("ENV_MYSELF", "foo")
	gossip := Gossip{Sender:"foo", Message:"bar"}
	jsonReq, err := json.Marshal(gossip)

	req, err := http.NewRequest("POST", "/whisper", bytes.NewBuffer(jsonReq))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ReceiveGossipHandler)

	//act
	handler.ServeHTTP(responseRecorder, req)

	//assert
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned failure status code: expect %v actual %v",
			http.StatusOK, status)
	}

	expectedBody := `that's myself`
	if responseRecorder.Body.String() != expectedBody {
		t.Error("expect handler to stop gossip but it didn't")
	}
}

func TestReceiveGossipHandler_WhenOthers_PassGossip(t *testing.T) {
	//setup
	os.Setenv("ENV_MYSELF", "foo")
	gossip := Gossip{Sender:"bar", Message:"bas"}
	jsonReq, err := json.Marshal(gossip)

	req, err := http.NewRequest("POST", "/whisper", bytes.NewBuffer(jsonReq))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ReceiveGossipHandler)

	// build a test http server to verify gossip is posted to the server
	// with the right URL, method, header, and body
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
    	w.WriteHeader(http.StatusOK)
    	if r.URL.EscapedPath() != "/whisper" {
      	t.Errorf("Expected request to '/whisper', got '%s'", r.URL.EscapedPath())
			}
			if r.Method != "POST" {
      	t.Errorf("Expected POST request, got '%s'", r.Method)
    	}
			if r.Header.Get("Content-type") != "application/json; charset=utf-8" {
				t.Errorf("Expected content-type 'application/json; charset=utf-8', got '%s'",
					 			 r.Header.Get("Content-type"))
			}
			var gossip Gossip
			json.NewDecoder(r.Body).Decode(&gossip)
    	if gossip.Sender != "bar" || gossip.Message != "bas" {
				 t.Errorf("Expected body 'sender:bar,message=bas', got: 'sender:%s,message:%s'",
				 				 gossip.Sender, gossip.Message)
    	}
  	}))
	defer ts.Close()
	
	os.Setenv("ENV_FORWARDURL", ts.URL + "/whisper")

	//act
	handler.ServeHTTP(responseRecorder, req)

	//assert
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned failure status code: expect %v actual %v",
			http.StatusOK, status)
	}

	if responseRecorder.Body.String() != "" {
		t.Error("expect handler to have empty body but it didn't")
	}
}