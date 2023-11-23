package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var store = sessions.NewCookieStore([]byte("your-very-secret-key")) // << TODO

func getHome(w http.ResponseWriter, r *http.Request) {

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

func handleLogin(s *Server, w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		log.Println("Authorization header missing or invalid")
		return
	}

	user := &User{Username: username, Password: password}

	log.Printf("User: %s, Pass: %s\n", user.Username, user.Password)

	// << TODO --------- VERIFY USER CREDS -----------
	// throw / return if invalid

	// Authentication successful, set session
	session, _ := store.Get(r, "chat-session")
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username

	err := session.Save(r, w)
	if err != nil {
		log.Printf("Error saving session: %v\n", err)
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Logged in successfully"))

}

func handleWs(s *Server, w http.ResponseWriter, r *http.Request) {

	// session, err := store.Get(r, "chat-session")
	// if err != nil {
	// 	http.Error(w, "Error fetching session", http.StatusInternalServerError)
	// 	return
	// }

	// // Check if the session has a username and it's a string
	// username, ok := session.Values["username"].(string)
	// if !ok || username == "" {
	// 	// Handle the case where the username is not set or not a string
	// 	log.Println("ERROR: username could not be found (handleWs)")
	// 	http.Error(w, "Unauthorized - No session username found", http.StatusUnauthorized)
	// 	return
	// }

	username := r.URL.Query().Get("username")
	if username == "" {
		log.Println("ERROR: username is empty")
		http.Error(w, "Unauthorized - No username provided", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {

		if r.Header.Get("Upgrade") != "websocket" {
			log.Println("Not a WebSocket upgrade request")
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{user: username, ws: conn, server: s, message: make(chan *Task)}
		client.server.register <- client

		// Spin client
		log.Printf("INFO: go client %s", username)
		go client.listen()
		go client.write()
	}
}

func testAuth(w http.ResponseWriter, r *http.Request) {

}

func main() {

	server := &Server{
		connections: make(map[string]*Client),
		tasks:       make(chan *Task),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}

	go server.spin()

	mux := http.NewServeMux()

	mux.Handle("/", logRequest(enableCors(http.HandlerFunc(getHome))))

	mux.Handle("/login", logRequest(enableCors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleLogin(server, w, r)
	}))))

	mux.Handle("/ws", logRequest(enableCors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleWs(server, w, r)
	}))))

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
