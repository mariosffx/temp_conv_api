package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9876"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/convert", corsMiddleware(handleConvert))

	log.Printf("TempConv REST server starting on :%s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
