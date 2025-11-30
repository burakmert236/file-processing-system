package nats_client

import (
	"log"

	utils "github.com/burakmert236/file-processing-system/internal/utils"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type NATS struct {
	NatsConnection   *nats.Conn
	JetStreamContext nats.JetStreamContext
}

type Subjects int

const (
	FileUplaoded Subjects = iota
	FileValidated
	FileValidationFailed
)

var subjectNames = map[Subjects]string{
	FileUplaoded:         "FileUploaded",
	FileValidated:        "FileValidated",
	FileValidationFailed: "FileValidationFailed",
}

func (s Subjects) String() string {
	return subjectNames[s]
}

func InitNats() *NATS {
	natsURL := utils.GetEnv("NATS_URL", true)

	nc, err := nats.Connect(natsURL,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2e9),
	)
	if err != nil {
		log.Fatalf("[NATS] Failed to connect: %v", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("[NATS] Failed to init JetStream: %v", err)
	}

	var subjectNamesArr = make([]string, 0, len(subjectNames))
	for _, v := range subjectNames {
		subjectNamesArr = append(subjectNamesArr, v)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "main-stream",
		Subjects: subjectNamesArr,
	})

	log.Println("NATS JetStream initialized")
	return &NATS{NatsConnection: nc, JetStreamContext: js}
}

func (n *NATS) CloseNATS() {
	n.NatsConnection.Drain()
}

func (n *NATS) Publish(subjectName string, event proto.Message) error {
	data, err := proto.Marshal(event)
	if err != nil {
		return err
	}

	var publishError = n.NatsConnection.Publish(subjectName, data)
	if publishError != nil {
		return publishError
	}

	return nil
}
