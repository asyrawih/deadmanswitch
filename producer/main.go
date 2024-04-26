package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func worker(id int, producer *kafka.Producer, topic string, wg *sync.WaitGroup, tasks <-chan int) {
	defer wg.Done()
	for task := range tasks {
		// Process the task
		time.Sleep(200 * time.Millisecond)
		value := fmt.Sprintf("value_%d", task/2)
		key := fmt.Sprintf("key_%d", task/2)
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(value),
			Key:   []byte(key),
		}, nil)
		fmt.Printf("Worker %d processed task %d\n", id, task)
	}
}

func main() {
	topic := "p_auto_order_cancel"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	defer producer.Flush(100)

	// Define the number of workers
	numWorkers := 50

	// Create a channel to send tasks to workers
	tasks := make(chan int, numWorkers*2)

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, producer, topic, &wg, tasks)
	}

	// Send tasks to workers
	for i := 0; i < 1000; i++ {
		tasks <- i
	}

	// Close the tasks channel to signal workers that no more tasks will be sent
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
}
