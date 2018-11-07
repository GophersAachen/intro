package main

import (
	"fmt"
	"net/http"
	"sync"
)

type ChatServer struct {
	Name string

	Messages      []string
	MessagesMutex sync.Mutex
}

func (s *ChatServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("request %v\n", req)
	s.MessagesMutex.Lock()
	defer s.MessagesMutex.Unlock()

	if req.URL.Path == "/api/messages" {
		// send messages
		res.WriteHeader(200)
		return
	} else if req.URL.Path == "/api/send" {
		// receive message, append to s.Message
		return
	}

	res.WriteHeader(500)
}

func main() {
	s := &ChatServer{
		Name: "Gophers Aachen",
	}

	mux := http.NewServeMux()
	mux.Handle("/api", s)
	mux.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe("127.0.0.1:3000", mux)
}
