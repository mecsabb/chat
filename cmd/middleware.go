package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// for localhost ease of use
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// If this is just a preflight request, respond with OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("------ New Request ------")
		log.Printf("Method: %s\n", r.Method)
		log.Printf("Path: %s\n", r.URL.Path)
		log.Println("Headers:")

		for name, headers := range r.Header {
			for _, h := range headers {
				log.Printf("%v: %v\n", name, h)
			}
		}

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

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// checks if the user has an active session
func checkAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "chat-session")
		if err != nil {
			http.Error(w, "Error fetching session", http.StatusInternalServerError)
			return
		}

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Unauthorized - Please login first", http.StatusUnauthorized)
			return
		}

		// User is authenticated, proceed with the request
		next.ServeHTTP(w, r)
	})
}
