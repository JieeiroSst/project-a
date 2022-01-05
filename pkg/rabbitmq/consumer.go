package rabbitmq

import (
	"context"
	"github.com/JieeiroSst/itjob/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
)

type consumerMQ struct {
	amqpConn *amqp.Connection
	cfg      *config.Config
}

type ConsumerMQ interface {
	CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, error)
	worker(ctx context.Context, messages <-chan amqp.Delivery, message func(ctx context.Context, deliveryBody []byte) error)
	StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string, queue func(ctx context.Context, deliveryBody []byte) error) error
}

func NewConsumer(amqpConn *amqp.Connection,cfg  *config.Config) ConsumerMQ {
	return &consumerMQ{
		amqpConn: amqpConn,
		cfg:cfg,
	}
}

func (c *consumerMQ) CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, error) {
	ch, err := c.amqpConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Error amqpConn.Channel")
	}
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
	}
	queue, err := ch.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueDeclare")
	}
	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchangeName,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueBind")
	}
	err = ch.Qos(
		prefetchCount,  // prefetch count
		prefetchSize,   // prefetch size
		prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error  ch.Qos")
	}
	return ch, nil
}

func (c *consumerMQ) worker(ctx context.Context, messages <-chan amqp.Delivery, queue func(ctx context.Context, deliveryBody []byte) error) {
	for delivery := range messages {
		span, ctx := opentracing.StartSpanFromContext(ctx, "consumer.worker")
		err := queue(ctx, delivery.Body)
		if err != nil {
			if err := delivery.Reject(false); err != nil {
				log.Println("Err delivery.Reject")
			}
		} else {
			err = delivery.Ack(false)
			if err != nil {
				log.Println("Failed to acknowledge delivery")
			}
		}
		span.Finish()
	}
}

func (c *consumerMQ) StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string, queue func(ctx context.Context, deliveryBody []byte) error) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch, err := c.CreateChannel(exchange, queueName, bindingKey, consumerTag)
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	defer ch.Close()
	deliveries, err := ch.Consume(
		queueName,
		consumerTag,
		consumeAutoAck,
		consumeExclusive,
		consumeNoLocal,
		consumeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Consume")
	}

	for i := 0; i < workerPoolSize; i++ {
		go c.worker(ctx, deliveries, queue)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}