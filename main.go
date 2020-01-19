package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// MsgResponse 返回消息结构
type MsgResponse struct {
	Text string `json:"text"`
}

func main() {
	// associate URLs requested to functions that handle requests
	http.HandleFunc("/api/msg", sendMessage)

	// start web server
	log.Println("Listening on http://localhost:9999/")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// basic handler for /hello request
func sendMessage(w http.ResponseWriter, r *http.Request) {

	// Fprint writes to a writer, in this case the http response
	response := MsgResponse{"Sever 的问候"}
	json, err := json.Marshal(response)
	if err != nil {
		log.Println("JSON Encoding Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}
