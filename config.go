package nsq

import "github.com/nsqio/go-nsq"

type Config struct {
	Address string
	Topic   string
	Channel string
	Message string
	Handler nsq.Handler
}
