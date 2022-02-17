package consume

import (
	"fmt"
	"log"
	"reflect"

	"github.com/yangyulong/queuemsg/limit"
	"github.com/yangyulong/queuemsg/queue"
)

type Consumer struct {
	Limit *limit.TokenBucket
}

var Consume *Consumer

func init() {
	Consume = &Consumer{
		Limit: limit.NewTokenBucket(),
	}
	Consume.Limit.SetCapacity(10)
	Consume.Limit.SetRate(1)
	go Consume.Limit.TimePutTokenToBucket(1)
}

func (con *Consumer) Consume(client *queue.Client, idx int) {
	go con.Sub(client, idx)
}

func (con *Consumer) Sub(client *queue.Client, idx int) {
	topic := fmt.Sprintf("topic:%v", idx)
	MessageChann, err := client.Subscribe(topic)
	if err != nil {
		log.Println(err)
	}
	for {
		if !con.Limit.Allow() {
			continue
		}
		Message := client.GetMessage(MessageChann)
		switch Message.(type) {
		case string:
			fmt.Printf("message type is:%v, body is %v", reflect.TypeOf(Message), reflect.ValueOf(Message))
		}
	}
}
