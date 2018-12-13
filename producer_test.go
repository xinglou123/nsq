package nsq

import (
	"testing"
)

func TestProducer(t *testing.T) {
	producer := NewProducer(&Config{Address: "127.0.0.1:4150", Topic: "test"})
	defer producer.Close()

	for a := 0; a < 100; a++ {
		producer.Publish("1234")
	}

}
