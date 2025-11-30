package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	utils "github.com/burakmert236/file-processing-system/internal/utils"
)

func HandleUpload(w http.ResponseWriter, req *http.Request) {
	uploadsFolder := utils.GetEnv("UPLOADS_FOLDER", true)

	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(w, "Unexpected error occurred "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	userId := req.PathValue("userId")
	fileName := strings.Split(header.Filename, ".")[0]
	extension := strings.Split(header.Filename, ".")[1]
	fileId := fmt.Sprintf("%s-%d.%s", fileName, time.Now().Nanosecond(), extension)
	folder := fmt.Sprintf("%s/%s", uploadsFolder, userId)
	tempPath := fmt.Sprintf("%s/%s", folder, fileId)

	storeError := utils.StoreFile(folder, fileId, file)
	if storeError != nil {
		http.Error(w, "Cannot store file: "+storeError.Error(), 500)
		return
	}

	publishError := PublishFileUploaded(fileId, userId, fileName, tempPath)
	if publishError != nil {
		http.Error(w, "Publish error: "+publishError.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(response)
}
