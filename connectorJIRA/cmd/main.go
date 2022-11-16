package main

import (
	"connectorJIRA/pkg/apiserver"
	"connectorJIRA/pkg/properties"
	"log"
	"os"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
