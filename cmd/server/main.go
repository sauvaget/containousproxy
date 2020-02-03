package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type server struct {
	config Config
}

func main() {
	err := runServer()
	if err != nil {
		log.Fatal(err)
	}
}

func runServer() {
	s := &server{}
	err := envconfig.Process("cproxy", &s.config)
	if err != nil {
		return err
	}
}
