package main

import (
	"github.com/gitKashish/golog/cmd"
	"github.com/gitKashish/golog/pkg/logger"
)

func main() {
	logger.Info("Starting GoLog")
	cmd.Execute()
}
