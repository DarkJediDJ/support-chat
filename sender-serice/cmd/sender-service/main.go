package main

import (
	"context"
	"log"

	server "github.com/DarkJediDJ/support-chat/sender-serice/api"
	"github.com/segmentio/kafka-go"
)

const addr = ":8080"

func main() {
	a := server.App{}

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "user-messages", 0)
	if err != nil {
		if err != nil {
			log.Printf("An Error Occured %v", err)
		}
	}

	a.New(conn)

	a.Run(addr)
}
