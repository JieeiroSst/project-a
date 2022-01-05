package rabbitmq

import (
	"github.com/JieeiroSst/itjob/config"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type publisherMQ struct {
	amqpChan *amqp.Channel
	cfg      *config.Config
}

type PublisherMQ interface {
	SetupExchangeAndQueue(exchange, queueName, bindingKey, consumerTag string) error
	CloseChan()
	Publish(body []byte, contentType string) error
}

func NewPublisher(cfg *config.Config) (PublisherMQ, error) {
	mqConn, err := NewRabbitMQConn(cfg)
	if err != nil {
		return nil, err
	}
	amqpChan, err := mqConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "p.amqpConn.Channel")
	}

	return &publisherMQ{
		cfg: cfg,
		amqpChan: amqpChan,
	}, nil
}

func (p *publisherMQ) SetupExchangeAndQueue(exchange, queueName, bindingKey, consumerTag string) error {
	err := p.amqpChan.ExchangeDeclare(
		exchange,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.ExchangeDeclare")
	}

	queue, err := p.amqpChan.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.QueueDeclare")
	}

	err = p.amqpChan.QueueBind(
		queue.Name,
		bindingKey,
		exchange,
		queueNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.QueueBind")
	}
	return nil
}

func (p *publisherMQ) CloseChan() {
	if err := p.amqpChan.Close(); err != nil {
		log.Printf("EmailsPublisher CloseChan: %v", err)
	}
}

func (p *publisherMQ) Publish(body []byte, contentType string) error {
	if err := p.amqpChan.Publish(
		p.cfg.RabbitMQ.Exchange,
		p.cfg.RabbitMQ.RoutingKey,
		publishMandatory,
		publishImmediate,
		amqp.Publishing{
			ContentType:  contentType,
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Body:         body,
		},
	); err != nil {
		return errors.Wrap(err, "ch.Publish")
	}
	return nil
}
