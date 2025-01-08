package broker

import (
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/broker/brokers/rabbitmq"
	"github.com/pmoura-dev/esr-service/internal/config"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Broker interface {
	GetSubscriber() message.Subscriber
	GetPublisher() message.Publisher

	Format(topic string) string

	Close()
}

func GetBroker(config config.BrokerConfig) (Broker, error) {
	switch config.BrokerType {
	case rabbitmq.Name:
		return rabbitmq.NewRabbitMQBroker(config)
	default:
		return nil, fmt.Errorf("unknown broker type: %s", config.BrokerType)
	}
}
