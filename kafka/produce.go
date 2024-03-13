package main

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// get kafka writer using environment variables.
	kafkaURL := "119.91.144.130:19092"
	topic := "just_a_topic"
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	go func() {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		for i := 0; i < 20; i++ {
			messages, err := conn.WriteMessages(
				kafka.Message{Value: []byte("one!")},
				kafka.Message{Value: []byte("two!")},
				kafka.Message{Value: []byte("three!")},
			)
			if err != nil {
				log.Fatalln(err)
			} else {
				log.Println("sent", messages)
			}
		}
	}()
	fmt.Println("wait")
	time.Sleep(time.Second * 100)
}
