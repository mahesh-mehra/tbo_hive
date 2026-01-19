package clients

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
)

var (
	ScyllaSession *gocql.Session
	RedisClient   *redis.Client
	KafkaProducer *kafka.Producer
)
