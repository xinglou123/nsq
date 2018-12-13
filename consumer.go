package nsq

import (
	"github.com/nsqio/go-nsq"
	"time"
)

type Consumer struct {
	conf *Config
	cons *nsq.Consumer
}

//初始化消费者
func NewConsumer(conf *Config) *Consumer {
	nsqconfig := nsq.NewConfig()
	nsqconfig.DefaultRequeueDelay = 0
	nsqconfig.MaxBackoffDuration = 20 * time.Millisecond
	nsqconfig.LookupdPollInterval = 1000 * time.Millisecond
	nsqconfig.RDYRedistributeInterval = 1000 * time.Millisecond
	nsqconfig.MaxInFlight = 2500

	c, err := nsq.NewConsumer(conf.Topic, conf.Channel, nsqconfig) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0) //屏蔽系统日志
	return &Consumer{
		conf: conf,
		cons: c,
	}
}
func (cs *Consumer) Handler(handler nsq.Handler) {
	if cs.cons != nil {
		cs.cons.AddHandler(handler) // 添加消费者接口
	}
}
func (cs *Consumer) Run() error {
	if cs.cons != nil {
		//建立NSQLookupd连接
		if err := cs.cons.ConnectToNSQD(cs.conf.Address); err != nil {
			return err
		}
	}
	return nil
}
func (cs *Consumer) Stop() {
	if cs.cons != nil {
		cs.cons.Stop()
	}
}
