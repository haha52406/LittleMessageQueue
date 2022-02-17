package queue

import (
	"errors"
	"sync"
)

type Broker interface {
	subscribe(topic string) (<-chan interface{}, error)
	unSubscribe(topic string, ch chan interface{}) error
	publish(topic string, msg interface{}) error
	setCapacity(capacity int)
}

type MessageQueue struct {
	Exit     chan interface{}
	Capacity int
	Topices  map[string][]chan interface{}

	sync.RWMutex
}

func (mq *MessageQueue) subscribe(topic string) (<-chan interface{}, error) {
	select {
	case <-mq.Exit:
		return nil, errors.New("mq subscribe is exit")
	default:
	}

	mq.Lock()
	defer mq.Unlock()
	chann := make(chan interface{})
	mq.Topices[topic] = append(mq.Topices[topic], chann)
	return chann, nil
}

func (mq *MessageQueue) unSubscribe(topic string, ch chan interface{}) error {
	select {
	case <-mq.Exit:
		return errors.New("mq unSubscribe is exit")
	default:
	}

	if _, ok := mq.Topices[topic]; !ok {
		return errors.New("mq unSubscribe topic is not exist")
	}

	mq.Lock()
	defer mq.Unlock()
	topics := mq.Topices[topic]
	newChann := make([]chan interface{}, 0)
	for _, chann := range topics {
		if chann == ch {
			continue
		}
		newChann = append(newChann, chann)
	}
	mq.Topices[topic] = newChann
	return nil
}

func (mq *MessageQueue) publish(topic string, msg interface{}) error {
	select {
	case <-mq.Exit:
		return errors.New("mq publish is exit")
	default:
	}

	if _, ok := mq.Topices[topic]; !ok {
		return errors.New("mq publish topis is not exist")
	}

	topics := mq.Topices[topic]
	for i := 0; i < len(topics); i++ {
		go func(idx int) {
			topics[idx] <- msg
		}(i)
	}
	return nil
}

func (mq *MessageQueue) setCapacity(capacity int) {
	mq.Capacity = capacity
}

func (mq *MessageQueue) getMessage(ch <-chan interface{}) interface{} {
	for message := range ch {
		if message != nil {
			return message
		}
	}
	return nil
}
