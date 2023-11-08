package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// for localhost ease of use
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("PING /\n")

	resp := make(map[string]string)
	resp["message"] = "Connection: Success"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func logRequest(r *http.Request) {
	log.Println("------ New Request ------")
	log.Printf("Method: %s\n", r.Method)
	log.Printf("Path: %s\n", r.URL.Path)
	log.Println("Headers:")

	for name, headers := range r.Header {
		for _, h := range headers {
			log.Printf("%v: %v\n", name, h)
		}
	}

	// Can remove if body of requests are long
	if r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			return
		}

		// Ensure the request body can be read again for further processing
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		log.Printf("Body: %s\n", string(bodyBytes))
	}

	log.Println("------ End of Request ------")
}

func handleLogin(s *Server, w http.ResponseWriter, r *http.Request) {

	logRequest(r)

	if r.Method == "GET" {

		if r.Header.Get("Upgrade") != "websocket" {
			log.Println("Not a WebSocket upgrade request")
			return
		}

		username, password, ok := r.BasicAuth()
		if !ok {
			log.Println("Authorization header missing or invalid")
			return
		}

		user := &User{Username: username, Password: password}

		log.Printf("User: %s, Pass: %s\n", user.Username, user.Password)

		// << TODO --------- VERIFY USER CREDS -----------
		// throw / return if invalid

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{user: user.Username, ws: conn, server: s, message: make(chan *Task)}
		client.server.register <- client

		// Spin client
		go client.listen()
		go client.write()
	}
}

func main() {

	server := &Server{
		connections: make(map[string]*Client),
		tasks:       make(chan *Task),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}

	go server.spin()

	http.HandleFunc("/", getHome)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handleLogin(server, w, r)
	})

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
