package main

import (
	"github.com/burakmert236/file-processing-system/file-validation/app"
	utils "github.com/burakmert236/file-processing-system/internal/utils"
)

func main() {
	app.Init()
	defer app.CloseNATS()
	utils.WaitForGracefulShutdown()
}
