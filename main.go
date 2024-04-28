package main

import (
	"kafka_compact/databases"
	"kafka_compact/deadman"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	topic := "p_auto_order_cancel"

	config := &kafka.ConfigMap{
		"bootstrap.servers":  "localhost:29092",
		"group.id":           "d_deadman",
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(config)

	if err != nil {
		slog.Error(err.Error())
	}

	defer c.Close()

	c.Subscribe(topic, nil)

	d := databases.NewDatabase()

	deadman := deadman.NewDeadMan(d)

	go func() {
		for {
			select {
			case <-sigchan:
				return
			default:
				e := c.Poll(100)
				switch evt := e.(type) {
				case *kafka.Message:
					// cek if have a message
					go d.Store(string(evt.Key), 2000)
					go deadman.Run(string(evt.Key))
					// go d.Reset(string(evt.Key), 100)

				case kafka.Error:
					return
				}
			}
		}
	}()

	<-sigchan

}
