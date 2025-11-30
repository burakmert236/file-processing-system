package app

import natsconn "github.com/burakmert236/file-processing-system/internal/nats_client"

var NATS *natsconn.NATS

func InitNATS() {
	NATS = natsconn.InitNats()
}

func CloseNATS() {
	NATS.CloseNATS()
}
