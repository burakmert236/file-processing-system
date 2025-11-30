package app

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	events "github.com/burakmert236/file-processing-system/generated/events"
)

func ValidateUploadedFile(event *events.FileUploaded) (bool, error) {
	info, err := os.Stat(event.TempPath)
	if err != nil {
		return false, errors.New("file not found on temporary storage")
	}

	const maxSize = 20 * 1024 * 1024
	if info.Size() > maxSize {
		return false, errors.New("file is too large (max 20 MB)")
	}

	extension := strings.ToLower(filepath.Ext(event.TempPath))

	allowed := map[string]bool{
		".pdf": true,
		".txt": true,
		".jpg": true,
		".png": true,
	}

	if !allowed[extension] {
		return false, errors.New("invalid file type: " + extension)
	}

	return true, nil
}
