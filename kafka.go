package api_kafka

import (
	"crypto/tls"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

var (
	RetryToConnect = time.Duration(10) * time.Second
	DialerTimeOut  = time.Duration(5) * time.Second
	ListenTimeout  = time.Duration(5) * time.Second
)

type Logger interface {
	Println(v ...interface{})
}

func NewDialer(username, password string) *kafka.Dialer {
	dialer := &kafka.Dialer{
		Timeout:   DialerTimeOut,
		DualStack: true,
	}
	if username != "kafka" && password != "kafka" {
		tls := &tls.Config{
			InsecureSkipVerify: true,
			ClientAuth:         0,
		}
		mechanism := plain.Mechanism{
			Username: username,
			Password: password,
		}
		dialer.SASLMechanism = mechanism
		dialer.TLS = tls
	}
	return dialer
}

func NewReader(username, password, groupID, topic string, brokers []string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		Dialer:         NewDialer(username, password),
		IsolationLevel: kafka.ReadUncommitted,
	})
}

func NewWriter(username, password, topic string, brokers []string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		Dialer:       NewDialer(username, password),
		RequiredAcks: 1,
	})
}
func NewOptions() *Options {
	return &Options{}
}
