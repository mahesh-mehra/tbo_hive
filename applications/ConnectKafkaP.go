package applications

import (
	"fmt"
	"os"
	"tbo_backend/clients"
	"tbo_backend/objects"
	"tbo_backend/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

// var deliveryChan = make(chan kafka.Event, 10000)

func ConnectKafkaP() bool {

	defer utils.HandlePanic()

	var err error

	fmt.Println("Initiating Kafka brokers connection on : ", objects.ConfigObj.KafkaBrokers)

	// getting hostname of the operating system
	name, err := os.Hostname()

	if err != nil {
		fmt.Println(err)
		return false
	}

	// creating a new kafka producer
	clients.KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        objects.ConfigObj.KafkaBrokers,
		"client.id":                name,
		"acks":                     "all",
		"socket.keepalive.enable":  "true",
		"message.send.max.retries": "100000",
		"queue.buffering.max.ms":   10,
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	// listening to kafka producer events in seperate go-routine
	go func() {

		defer utils.HandlePanic()

		for e := range clients.KafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	fmt.Println("Kafka brokers connected on : ", objects.ConfigObj.KafkaBrokers)

	return true
}

// GetKafkaP returning kafka producer object
func GetKafkaP() *kafka.Producer {
	defer utils.HandlePanic()
	return clients.KafkaProducer
}

// KafkaPublish method to publish messages to kafka topic
func KafkaPublish(topic string, key *string, data string) {

	defer utils.HandlePanic()

connectKafka:
	var kafka_key string

	// if key is null then setting random UUID, it will push message method in a round robin way
	// if there is multiple partitions in a single topic
	if key == nil {
		kafka_key = uuid.New().String()
	} else {
		kafka_key = *key
	}

	// if kafka producer is not null then publishing messages to kafka topics in asynchronous way
	// using golang channels
	if clients.KafkaProducer != nil {

		err := clients.KafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(kafka_key),
			Value:          []byte(data),
		}, nil)

		if err != nil {
			fmt.Println(err)
			clients.KafkaProducer.Close()
			clients.KafkaProducer = nil
			// ConnectKafkaP()
			clients.KafkaProducer.Flush(10)
			goto connectKafka
		}

		clients.KafkaProducer.Flush(10)
	}
}
