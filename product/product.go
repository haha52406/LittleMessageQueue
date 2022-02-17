package product

import (
	"fmt"
	"log"
	"time"

	"github.com/yangyulong/queuemsg/queue"
)

func Producter(client *queue.Client) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < 10; i++ {
				topic := fmt.Sprintf("topic:%v", i)
				msg := fmt.Sprintf("producter:%v\n", i)
				fmt.Printf("i am producter:%v, msg:%v", i, msg)
				err := client.Publish(topic, msg)
				if err != nil {
					log.Println(err)
				}
			}
		default:
		}
	}

}
