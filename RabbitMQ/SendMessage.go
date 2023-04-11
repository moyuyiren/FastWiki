package RabbitMQ

import (
	"github.com/streadway/amqp"
)

func SendMessage(str string) error {
	err := ch.Publish("sms_server", "sms_send", false, false,
		amqp.Publishing{ContentType: "text/plain",
			Body: []byte(str),
		})
	if err != nil {
		return err
	}
	return nil
}
