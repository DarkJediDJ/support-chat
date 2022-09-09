package main

import (
	"context"
	"log"

	server "github.com/DarkJediDJ/support-chat/sender-serice/api"
	"github.com/segmentio/kafka-go"
)

const addr = ":8080"

func main() {
	a := server.Init()

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "user-messages", 0)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer conn.Close()

	a.InitRouter(conn)

	a.Run(addr)
}
