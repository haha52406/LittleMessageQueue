package queue

type Client struct {
	Bro *MessageQueue
}

func NewClient() *Client {
	return &Client{
		Bro: &MessageQueue{
			Exit:     make(chan interface{}),
			Capacity: 0,
			Topices:  make(map[string][]chan interface{}),
		},
	}
}
func (client *Client) Subscribe(topic string) (<-chan interface{}, error) {
	return client.Bro.subscribe(topic)
}

func (client *Client) UnSubscribe(topic string, ch chan interface{}) error {
	return client.Bro.unSubscribe(topic, ch)
}

func (client *Client) SetCapacity(capacity int) {
	client.Bro.setCapacity(capacity)
}

func (client *Client) Publish(topic string, msg interface{}) error {
	return client.Bro.publish(topic, msg)
}

func (client *Client) GetMessage(ch <-chan interface{}) interface{} {
	return client.Bro.getMessage(ch)
}
