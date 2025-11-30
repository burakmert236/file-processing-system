package app

import (
	events "github.com/burakmert236/file-processing-system/generated/events"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func FileUploadHandler(message *nats.Msg) {
	var event events.FileUploaded
	proto.Unmarshal(message.Data, &event)
	result, error := ValidateUploadedFile(&event)

	if result {
		PublishFileValidated(event.File)
	} else {
		PublishFileValidationFailed(event.File, error.Error())
	}
}
