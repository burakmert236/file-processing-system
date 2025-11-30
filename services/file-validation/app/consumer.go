package app

import (
	"log"

	"github.com/burakmert236/file-processing-system/internal/nats_client"
	"github.com/nats-io/nats.go"
)

func StartConsumer() {
	onMessage := func(message *nats.Msg) {
		log.Printf("Message received: %s", message.Data)
		FileUploadHandler(message)
		message.Ack()
	}

	NATS.JetStreamContext.Subscribe(nats_client.FileUplaoded.String(), onMessage, nats.Durable("file-validation"), nats.ManualAck())
}
