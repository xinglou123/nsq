package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"testing"
)

// 消费者
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func TestConsumer(t *testing.T) {
	consumer := NewConsumer(&Config{Address: "127.0.0.1:4150", Topic: "test", Channel: "logger-channel"})
	defer consumer.Stop()
	consumer.Handler(&ConsumerT{})
	consumer.Run()
	select {}
}
