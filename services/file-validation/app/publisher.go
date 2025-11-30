package app

import (
	"log"

	"github.com/burakmert236/file-processing-system/generated/common"
	events "github.com/burakmert236/file-processing-system/generated/events"
	"github.com/burakmert236/file-processing-system/internal/nats_client"
)

func PublishFileValidated(file *common.FileRef) error {
	event := &events.FileValidated{
		File: file,
	}

	log.Printf("Message sent: %s", event)
	return NATS.Publish(nats_client.FileValidated.String(), event)
}

func PublishFileValidationFailed(file *common.FileRef, message string) error {
	event := &events.FileValidationFailed{
		File:    file,
		Message: message,
	}

	log.Printf("Message sent: %s", event)
	return NATS.Publish(nats_client.FileValidationFailed.String(), event)
}
