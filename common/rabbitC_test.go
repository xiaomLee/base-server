package common

import (
	"fmt"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

func TestInitRabbit(t *testing.T) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s%s", "dev", "dev", "127.0.0.1", "5672", "/dev")
	if err := InitRabbit(url); err != nil {
		t.Fatal(err)
	}
	rabbit := GetRabbitInstance()
	rabbit.ExchangeDeclare("dev_exchange")
	rabbit.QueueDeclare("info", "dev_exchange", "info")
	rabbit.QueueDeclare("error", "dev_exchange", "error")
	rabbit.QueueDeclare("warning", "dev_exchange", "warning")

	rabbit.Consume("info_log_1", "info", infoLog)
	rabbit.Consume("info_log_2", "info", infoLog)
	rabbit.Consume("error_log", "error", errorLog)

	ticker := time.NewTicker(1 * time.Second)
	timer := time.After(100 * time.Second)
	for {
		select {
		case <-ticker.C:
			str := time.Now().String()
			rabbit.PushTransientMessage("dev_exchange", "info", []byte(str))
			rabbit.PushTransientMessage("dev_exchange", "error", []byte(str))
		case <-timer:
			return
			//println("restart success")
		}
	}
}

func infoLog(d amqp.Delivery) {
	fmt.Printf("info consumer[%s] msg:%s \n", d.ConsumerTag, string(d.Body))
}

func errorLog(d amqp.Delivery) {
	fmt.Printf("error consumer[%s] msg:%s \n", d.ConsumerTag, string(d.Body))
}
