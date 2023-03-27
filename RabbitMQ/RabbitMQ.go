package RabbitMQ

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var Conn *amqp.Connection
var Ch *amqp.Channel

func InitRabbitMq(err error) {
	Conn, err := amqp.Dial("amqp://yiren:asx123..@120.77.42.228:5672/")
	if err != nil {
		zap.L().Error("RabbitMQ Connection Field", zap.Error(err))
		return
	}
	Ch, err := Conn.Channel()
	if err != nil {
		zap.L().Error("RabbitMQ Create Channel Field", zap.Error(err))
		return
	}
	err = Ch.ExchangeDeclare("")

}

func CloseRabbitMqConn() {
	Conn.Close()
	Ch.Close()
	return
}
