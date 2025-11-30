package nats_client

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type NATS struct {
	NatsConnection   *nats.Conn
	JetStreamContext nats.JetStreamContext
}

func InitNats(subjectNames []string) *NATS {
	natsURL := os.Getenv("NATS_URL")

	if natsURL == "" {
		log.Fatal("NATS_URL env variable is required")
	}

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

	js.AddStream(&nats.StreamConfig{
		Name:     "main-stream",
		Subjects: subjectNames,
	})

	log.Println("NATS JetStream initialized")
	return &NATS{NatsConnection: nc, JetStreamContext: js}
}

func (n *NATS) CloseNATS() {
	n.NatsConnection.Drain()
}
