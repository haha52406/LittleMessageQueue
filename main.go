package main

import (
	"github.com/yangyulong/queuemsg/consume"
	"github.com/yangyulong/queuemsg/product"
	"github.com/yangyulong/queuemsg/queue"
)

func main() {
	client := queue.NewClient()
	for i := 0; i < 10; i++ {
		consume.Consume.Consume(client, i)
	}
	product.Producter(client)
}
