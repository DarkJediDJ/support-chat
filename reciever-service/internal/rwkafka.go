package internal

import (
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type IncMessage struct {
	Id     int    `json:"ID"`
	Text   string `json:"message"`
	Answer string `json:"answer,omitempty"`
}

func KafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func KafkaWriter(conn *kafka.Conn, body string) {
	err := conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	_, err = conn.Write([]byte(body))
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
}
