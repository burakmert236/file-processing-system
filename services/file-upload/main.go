package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/burakmert236/file-processing-system/file-upload/app"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT env variable is required")
	}

	app.InitNATS()
	defer app.CloseNATS()

	http.HandleFunc("/upload/user/{userId}", app.HandleUpload)

	log.Printf("File Upload service running on :%s", port)
	go http.ListenAndServe(":"+port, nil)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down...")
}
