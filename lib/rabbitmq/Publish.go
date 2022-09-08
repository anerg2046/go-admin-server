package rabbitmq

import (
	"go-app/lib/logger"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// 直接模式,生产者.
func (r *RabbitMQ) PublishSimple(message []byte) (err error) {

	//第一步，申请队列，如不存在，则自动创建之，存在，则路过。
	_, err = r.channel.QueueDeclare(
		r.config.QueueName,
		r.config.Durable,
		r.config.AutoDelete,
		r.config.Exclusive,
		r.config.NoWait,
		r.config.Args,
	)
	if err != nil {
		logger.Error(LOG_FIELD, zap.Error(err))
		return
	}

	//第二步，发送消息到队列中
	err = r.channel.Publish(
		r.config.ExChange,
		r.config.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		logger.Error(LOG_FIELD, zap.Error(err))
	}
	return
}
