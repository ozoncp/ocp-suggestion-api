package producer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"

	cfg "github.com/ozoncp/ocp-suggestion-api/internal/config"
)

type Producer interface {
	Send(topic string, message *Message) error
	Close() error
}

type producer struct {
	prod sarama.SyncProducer
}

// NewProducer возвращает структуру producer (интерфейс Producer)
func NewProducer() (*producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.Config.Brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		prod: syncProducer,
	}, nil
}

// Send отправляет сообщение в брокер
func (p *producer) Send(topic string, message *Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed marshaling message to json: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(bytes),
		Timestamp: time.Time{},
	}
	_, _, err = p.prod.SendMessage(msg)
	return err
}

// Close закрывает соединение с брокером
func (p *producer) Close() error {
	return p.prod.Close()
}
