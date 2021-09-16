package main

import (
	"github.com/marcoshuck/logger"
	"os"
)

func main() {
	log := logger.NewLogger(logger.VerbosityInfo, os.Stdout)
	log.Debug("Hello, world!")
}
