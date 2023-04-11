package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"testing"
)

func Test_Reply(t *testing.T) {
	conn, _ := amqp.Dial("amqp://root:root@127.0.0.1:5672/")
	ch, _ := conn.Channel()

	//监听queue.dlx队列
	msgs, _ := ch.Consume("sms", "", true, false, false, false, nil)

	for d := range msgs {
		fmt.Printf("receive: %s\n", d.Body) // 收到消息，业务处理
	}
}
