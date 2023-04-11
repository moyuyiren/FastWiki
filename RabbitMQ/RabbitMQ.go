package RabbitMQ

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var ch *amqp.Channel
var conn *amqp.Connection

func InitRabbitMq() error {
	conn, err := amqp.Dial("amqp://root:root@127.0.0.1:5672/")
	if err != nil {
		zap.L().Error("RabbitMQ Connection Field", zap.Error(err))
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		zap.L().Error("RabbitMQ Create Channel Field", zap.Error(err))
		return err
	}
	err = ch.ExchangeDeclare(
		"sms_server", // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		zap.L().Error("RabbitMQ Create Exchange Field", zap.Error(err))
		return err
	}

	q, err := ch.QueueDeclare(
		"sms", // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		zap.L().Error("RabbitMQ Create Queue Field", zap.Error(err))
		return err
	}
	err = ch.QueueBind(
		q.Name,       // queue name
		"sms_send",   // routing key
		"sms_server", // exchange
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("RabbitMQ Bind Queue Field", zap.Error(err))
		return err
	}

	return nil

}

func Close() {
	conn.Close()
	ch.Close()
	return
}
