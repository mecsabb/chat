package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// dependency injection
type application struct {
}

// for localhost ease of use
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	//router := httprouter.New()

	http.HandleFunc("/", getHome)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
