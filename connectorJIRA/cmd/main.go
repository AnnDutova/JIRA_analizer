package main

import (
	"connectorJIRA/pkg/apiserver"
	"connectorJIRA/pkg/properties"
	"log"
	"os"
)

func main() {
	config := properties.GetConfig(os.Args[1])

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
