package main

import (
	"github.com/joho/godotenv"
	"github.com/peake100/stalkforecast-go/rest"
	"github.com/peake100/stalkforecast-go/service"
	"log"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)

	serviceErrs := make(chan error)
	restErrs := make(chan error)

	go service.RunService(serviceErrs)
	go rest.RunRest(restErrs)

	select {
	case err := <- serviceErrs:
		log.Fatal(err)
	case err := <- restErrs:
		log.Fatal(err)
	}
}
