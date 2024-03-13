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
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		batch := conn.ReadBatch(10e3, 1e6)
		bytes := make([]byte, 10e3)
		for {
			_, err := batch.Read(bytes)
			if err != nil {
				break
			}
			log.Println("received", string(bytes))
		}
		batch.Close()
	}()
	fmt.Println("wait")
	time.Sleep(time.Second * 100)
}
