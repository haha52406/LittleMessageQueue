package limit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	TokenNum int
	Capacity int
	Rate     int

	sync.Mutex
}

func NewTokenBucket() *TokenBucket {
	return &TokenBucket{
		TokenNum: 0,
		Capacity: 0,
		Rate:     0,
	}
}

func (tb *TokenBucket) SetCapacity(capacity int) {
	tb.Capacity = capacity
}

func (tb *TokenBucket) SetRate(rate int) {
	tb.Rate = rate
}

func (tb *TokenBucket) TimePutTokenToBucket(putTokenRate int) {
	ticker := time.NewTicker(time.Duration(putTokenRate) * time.Second)
	for {
		select {
		case <-ticker.C:
			tb.putTokenToBucket()
		default:
		}
	}
}

func (tb *TokenBucket) putTokenToBucket() {
	tb.Lock()
	defer tb.Unlock()
	if tb.Capacity == tb.TokenNum {
		return
	}
	tb.TokenNum++
}

func (tb *TokenBucket) Allow() bool {
	tb.Lock()
	defer tb.Unlock()

	if tb.TokenNum > 0 {
		tb.TokenNum--
		return true
	} else {
		return false
	}
}
