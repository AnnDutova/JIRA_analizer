package main

import (
	"connectorJIRA/pkg/apiserver"
	"connectorJIRA/pkg/logging"
	"os"
)

func main() {
	logging.Init(os.Args[1])
	logger := logging.GetLogger()
	logger.Info("Start connector")

	if err := apiserver.Start(); err != nil {
		logger.Fatal(err)
	}
}
