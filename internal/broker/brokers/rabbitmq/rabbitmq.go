package rabbitmq

import (
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	Name = "rabbitmq"

	exchangeName = "esr-service"
	queueName    = "esr-service-queue"
)

type Broker struct {
	subscriber *amqp.Subscriber
	publisher  *amqp.Publisher
}

func NewRabbitMQBroker(config config.BrokerConfig) (*Broker, error) {
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)

	amqpConfig := amqp.NewDurableTopicConfig(amqpURI, exchangeName, queueName)

	subscriber, err := amqp.NewSubscriber(amqpConfig, watermill.NewSlogLogger(nil))
	if err != nil {
		return nil, err
	}

	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewSlogLogger(nil))
	if err != nil {
		return nil, err
	}

	return &Broker{
		subscriber: subscriber,
		publisher:  publisher,
	}, nil
}

func (b *Broker) GetSubscriber() message.Subscriber {
	return b.subscriber
}

func (b *Broker) GetPublisher() message.Publisher {
	return b.publisher
}

func (b *Broker) Close() {
	_ = b.publisher.Close()
	_ = b.subscriber.Close()
}
