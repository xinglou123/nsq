package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type Producer struct {
	conf *Config
	prod *nsq.Producer
}

// 初始化生产者
func NewProducer(conf *Config) *Producer {
	var err error
	nsqconfig := nsq.NewConfig()
	nsqconfig.DefaultRequeueDelay = 0
	nsqconfig.MaxBackoffDuration = 20 * time.Millisecond
	nsqconfig.LookupdPollInterval = 1000 * time.Millisecond
	nsqconfig.RDYRedistributeInterval = 1000 * time.Millisecond
	nsqconfig.MaxInFlight = 2500

	mprod, err := nsq.NewProducer(conf.Address, nsqconfig)
	fmt.Println(mprod)
	if err != nil {
		return nil
	}
	return &Producer{
		conf: conf,
		prod: mprod,
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
	return err
}

//
func (p *Producer) Close() {
	if p.prod != nil {
		p.prod.Stop()
	}
}
