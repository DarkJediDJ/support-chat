package main

import (
	"context"
	"encoding/json"
	"log"

	i "github.com/DarkJediDJ/support-chat/reciever-service/internal"
	"github.com/segmentio/kafka-go"
)

func main() {

	reader := i.KafkaReader("localhost:9092", "user-messages")
	defer reader.Close()

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "answer-messages", 0)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer conn.Close()

	var message i.IncMessage

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("An Error Occured %v", err)
			continue
		}

		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
			continue
		}

		message.Answer = "hi,aboba"

		outMesssage, err := json.Marshal(message)
		if err != nil {
			log.Printf("An Error Occured %v", err)
			continue
		}

		i.KafkaWriter(conn, string(outMesssage))
	}

}
