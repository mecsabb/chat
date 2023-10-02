package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	username string
	password string
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

func handleLogin(s *Server, w http.ResponseWriter, r *http.Request) {
	// Generic GET response for debug
	if r.Method == "GET" {
		resp := make(map[string]string)
		resp["GET"] = "/login Request: Success"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	}

	// postLogin is responsible for handing user login functionality
	if r.Method == "POST" {
		user := &User{}

		// Recieve user info in post request (decode JSON into user)
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("User: %s, Pass: %s", user.username, user.password)
		w.Write([]byte("Pre Upgrader"))
		// << TODO --------- VERIFY USER CREDS -----------
		// throw / return if invalid

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{user: user.username, ws: conn, server: s, message: make(chan *Task)}
		client.server.register <- client

		w.Write([]byte(client.user))

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
