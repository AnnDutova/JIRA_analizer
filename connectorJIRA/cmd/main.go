package main

import (
	"connectorJIRA/pkg/apiserver"
	"connectorJIRA/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("Start connector")

	if err := apiserver.Start(); err != nil {
		logger.Fatal(err)
	}
}
