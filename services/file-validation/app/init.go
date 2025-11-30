package app

import (
	"log"
)

func Init() {
	InitNATS()
	StartConsumer()

	log.Printf("File Validation service running")
}
