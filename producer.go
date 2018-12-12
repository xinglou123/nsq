package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Producer struct {
	conf *Config
	prod *nsq.Producer
}

// 初始化生产者
func NewProducer(conf *Config) *Producer {
	var err error
	producer, err := nsq.NewProducer(conf.Address, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	return &Producer{
		conf: conf,
		prod: producer,
	}
}

//发布消息
func (p *Producer) Publish(message string) error {
	var err error
	if p.prod != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = p.prod.Publish(p.conf.Topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil", err)
}

//
func (p *Producer) Close() {
	if p.prod != nil {
		p.prod.Stop()
	}
}
