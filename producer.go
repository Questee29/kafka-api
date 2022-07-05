package api_kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type Writer interface {
	Close() error
	WriteMessages(ctx context.Context, msg ...kafka.Message) error
}

type Producer struct {
	writer        Writer
	writerBuilder func() Writer
	log           Logger
}

func NewProducer(writerBuilder func() Writer, log Logger) *Producer {
	return &Producer{
		writer:        writerBuilder(),
		writerBuilder: writerBuilder,
		log:           log,
	}
}
func (queue *Producer) Publish(ctx context.Context, payloads ...interface{}) error {
	queue.writer = queue.writerBuilder()
	messages := make([]kafka.Message, len(payloads))
	for i, payload := range payloads {
		bytes, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		messages[i] = kafka.Message{
			Value: bytes,
		}
	}
	defer func() {
		if err := queue.writer.Close(); err != nil {
			queue.log.Println(err, "at publish")
		}
	}()
	return queue.writer.WriteMessages(ctx, messages...)
}
