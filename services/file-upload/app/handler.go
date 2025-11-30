package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/burakmert236/file-processing-system/generated/common"
	events "github.com/burakmert236/file-processing-system/generated/events"
)

func HandleUpload(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("file")

	if err != nil {
		http.Error(w, "Unexpected error occurred "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	userId := req.PathValue("userId")
	fileName := strings.Split(header.Filename, "")[0]
	fileId := fmt.Sprintf("%s-%d", fileName, time.Now().Nanosecond())

	event := &events.FileUploaded{
		File: &common.FileRef{
			FileId:   fileId,
			UserId:   userId,
			FileName: fileName,
		},
		TempPath: fmt.Sprintf("/tmp/%s/fileId", userId),
	}

	data, err := proto.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	NATS.NatsConnection.Publish(subjectName, data)
}
