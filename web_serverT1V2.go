package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)


type  RequestInfo struct {
	Host string `json:"host"`
	User_agent string `json:"user_agent"`
	Request_uri string `json:"request_uri"`
	Headers  http.Header`json:"headers"`
}
func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/Request", RequestHandler)
	s := http.Server{
		Addr:           ":8087",
		Handler:        handler,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		IdleTimeout:    100 * time.Second,
		MaxHeaderBytes: 1 << 40,
	}
	log.Fatal(s.ListenAndServe())
}

func RequestHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	request := RequestInfo{
		Host:r.Host,
		User_agent:r.UserAgent(),
		Request_uri:r.RequestURI,
		Headers:r.Header,
		}

	requestJson, _ := json.Marshal(request)
	w.WriteHeader(http.StatusOK)
	w.Write(requestJson)
}

