package main

import (
	"bytes"
	"context"
	"log"

	i "github.com/DarkJediDJ/support-chat/reciever-service/internal"
	"github.com/segmentio/kafka-go"
)

func main() {

	reader := i.KafkaReader("localhost:9092", "user-messages")
	defer reader.Close()

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "answer-messages", 0)
	if err != nil {
		if err != nil {
			log.Printf("An Error Occured %v", err)
		}
	}
	defer conn.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("An Error Occured %v", err)
		}
		ins := []byte(`, "answer": "hi, aboba!"`) // Note the leading comma.
		closingBraceIdx := bytes.LastIndexByte(m.Value, '}')
		m.Value = append(m.Value[:closingBraceIdx], ins...)
		m.Value = append(m.Value, '}')
		i.KafkaWriter(conn, string(m.Value))
	}

}
