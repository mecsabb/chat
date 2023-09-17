package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
