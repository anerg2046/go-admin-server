package rabbitmq

import (
	"fmt"
	"go-app/lib/logger"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (r *RabbitMQ) ConsumeSimple() {

	fmt.Printf("%+v\n", r)

	defer func() {
		if err := recover(); err != nil {
			logger.Suger().Info(LOG_FIELD, "3秒后重试")
			time.Sleep(3 * time.Second)
			r.Connect()
			r.ConsumeSimple()
		}
	}()

	//第一步,申请队列,如果队列不存在则自动创建,存在则跳过
	q, err := r.channel.QueueDeclare(
		r.config.QueueName,
		r.config.Durable,
		r.config.AutoDelete,
		r.config.Exclusive,
		r.config.NoWait,
		r.config.Args,
	)
	if err != nil {
		logger.Error(LOG_FIELD, zap.Error(err))
	}
	closeChan := make(chan *amqp.Error, 1)
	notifyClose := r.channel.NotifyClose(closeChan) //一旦消费者的channel有错误，产生一个amqp.Error，channel监听并捕捉到这个错误
	//第二步,接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		r.config.Consumer,
		r.config.AutoAck,
		r.config.Exclusive,
		r.config.NoLocal,
		r.config.NoWait,
		r.config.Args,
	)
	if err != nil {
		logger.Error(LOG_FIELD, zap.Error(err))
	}

LOOP:
	for {
		select {
		case e := <-notifyClose:
			logger.Error(LOG_FIELD, zap.Error(e))
			close(closeChan)
			break LOOP
		case msg := <-msgs:
			r.Payload <- msg.Body
		}
	}
}
