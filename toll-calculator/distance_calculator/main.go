package main

import (
	"log"

	"github.com/tunangoo/full-time-go-dev/toll-calculator/aggregator/client"
)

const (
	kafkaTopic         = "obuData"
	aggregatorEndpoint = "http://localhost:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc, err = NewCalculatorService()
	if err != nil {
		log.Fatal(err)
	}

	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
