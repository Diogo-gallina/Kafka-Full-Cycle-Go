package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main()  {
  producer := NewKafkaProducer()
}

func NewKafkaProducer() *kafka.Producer {
  configMap := &kafka.ConfigMap{
	"bootstrap.servers": "kafka-full-cycle-go_kafka_1:9092",
  }
  p, err := kafka.NewProducer(configMap)
  if err != nil {
	log.Println(err.Error())
  }
  return p
}
