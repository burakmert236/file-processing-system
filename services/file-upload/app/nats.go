package app

import natsconn "github.com/burakmert236/file-processing-system/internal/nats_client"

var NATS *natsconn.NATS

const subjectName string = "file-uploads"

func InitNATS() {
	NATS = natsconn.InitNats([]string{subjectName})
}

func CloseNATS() {
	NATS.CloseNATS()
}
