package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message": "Hello, World API!",
	}

	// Correctly encode response as JSON
	json.NewEncoder(w).Encode(response)
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello-world", HelloWorld).Methods("GET")
	return router
}

func main() {
	router := Router()
	fmt.Println("✅ Backend running on: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
