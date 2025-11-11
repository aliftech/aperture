package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello from User Service!", "path": "` + r.URL.Path + `"}`))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("User service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
