package RabbitMQ

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var Ch *amqp.Channel
var Conn *amqp.Connection

func InitRabbitMq() error {
	Conn, err := amqp.Dial("amqp://yiren:asx123..@120.77.42.228:5672/")
	if err != nil {
		zap.L().Error("RabbitMQ Connection Field", zap.Error(err))
		return err
	}
	Ch, err := Conn.Channel()
	if err != nil {
		zap.L().Error("RabbitMQ Create Channel Field", zap.Error(err))
		return err
	}
	err = Ch.ExchangeDeclare(
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

	q, err := Ch.QueueDeclare(
		"",    // name
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
	err = Ch.QueueBind(
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
	Conn.Close()
	Ch.Close()
	return
}
