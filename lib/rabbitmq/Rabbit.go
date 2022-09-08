package rabbitmq

import (
	"go-app/lib/logger"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

const LOG_FIELD = "[RABBITMQ]"

// RabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  RabbitConfig
	Payload chan []byte //拿到的数据

}

type RabbitConfig struct {
	Dsn        string     //队列地址
	QueueName  string     //队列名
	ExChange   string     //交换机名称
	Key        string     //绑定的key名称
	Durable    bool       //是否持久化
	AutoDelete bool       //是否自动删除
	Exclusive  bool       //是否具有排他性
	NoWait     bool       //是否阻塞处理
	Args       amqp.Table //额外的属性

	Consumer string //指定消费者
	AutoAck  bool   //是否自动应答,告诉我已经消费完了
	NoLocal  bool   //若设置为true,则表示为不能将同一个connection中发送的消息传递给这个connection中的消费者
}

// 创建结构体实例，参数:地址、队列名称、交换机名称和bind的key
func NewRabbitMQ(config RabbitConfig) *RabbitMQ {
	return &RabbitMQ{config: config, Payload: make(chan []byte, 512)}
}

func NewSimple(config RabbitConfig) (rabbitmq *RabbitMQ) {

	rabbitmq = NewRabbitMQ(config)
	rabbitmq.Connect()
	return rabbitmq
}

func (r *RabbitMQ) Connect() {
	defer func() {
		if err := recover(); err != nil {
			logger.Suger().Info(LOG_FIELD, "3秒后重试")
			time.Sleep(3 * time.Second)
			r.Connect()
		}
	}()

	var err error
	//获取参数connection
	r.conn, err = amqp.Dial(r.config.Dsn)
	if err != nil {
		logger.Error(LOG_FIELD, zap.Error(err))
	}
	//获取channel参数
	r.channel, err = r.conn.Channel()
	if err != nil {
		logger.Error(LOG_FIELD, zap.Error(err))
	}

}

// 关闭conn和chanel的方法
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}
