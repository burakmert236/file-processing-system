package app

import (
	"log"
	"net/http"

	utils "github.com/burakmert236/file-processing-system/internal/utils"
)

func Init() {
	port := utils.GetEnv("PORT", true)

	InitNATS()

	http.HandleFunc("/upload/user/{userId}", HandleUpload)

	log.Printf("File Upload service running on :%s", port)
	go http.ListenAndServe(":"+port, nil)
}
