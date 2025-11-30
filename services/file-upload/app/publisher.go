package app

import (
	"github.com/burakmert236/file-processing-system/generated/common"
	events "github.com/burakmert236/file-processing-system/generated/events"
	"github.com/burakmert236/file-processing-system/internal/nats_client"
)

func PublishFileUploaded(
	fileId string, userId string, fileName string, tempPath string,
) error {
	event := &events.FileUploaded{
		File: &common.FileRef{
			FileId:   fileId,
			UserId:   userId,
			FileName: fileName,
		},
		TempPath: tempPath,
	}

	return NATS.Publish(nats_client.FileUplaoded.String(), event)
}
