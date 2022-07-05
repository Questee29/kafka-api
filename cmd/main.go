package main

import (
	"context"
	"fmt"
	"log"

	kafkaq "github.com/Questee29/api_kafka"
	config "github.com/Questee29/api_kafka/configs"
)

func main() {
	cfg, err := config.LoadConfig("api", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(cfg.Consumer.KAFKA_TOPIC)
	opts := kafkaq.NewOptions()
	consumer := kafkaq.NewListener(
		func() kafkaq.Reader {
			return kafkaq.NewReader(cfg.Consumer.KAFKA_USER,
				cfg.Consumer.KAFKA_PASS,
				cfg.Consumer.NAME,
				cfg.Consumer.KAFKA_TOPIC,
				cfg.Consumer.KAFKA_BROKERS)
		},
		log.Default(),
		opts,
	)

	go func() {
		if err := consumer.Listen(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		msg := <-consumer.Msg()
		log.Println(string(msg))
	}
}
