package stream_util

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaImpl struct {
	BootstrapServers string
  UniqueConsumerID string
}

func (u KafkaImpl) Send(topic string, message []byte) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(u.BootstrapServers),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})

	return err
}

func (u KafkaImpl) Subscribe(topic string) (chan []byte, chan error, chan bool) {

	msgChannel := make(chan []byte)
	errChannel := make(chan error)
	unsubChannel := make(chan bool)

	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{u.BootstrapServers},
			Topic:    topic,
			MaxWait:  time.Second,
      GroupID: u.UniqueConsumerID,
		})

		defer reader.Close()

		for {
			select {
			case <-unsubChannel:
				return
			default:
				msg, err := reader.ReadMessage(context.Background())
				if err != nil {
					errChannel <- err
					return
				} else {
					msgChannel <- msg.Value
				}
			}
		}
	}()

	return msgChannel, errChannel, unsubChannel

}
