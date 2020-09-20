package observer_test

import (
	"github.com/nikitych1w/learn-go-patterns/Behavioral/observer"
	"sync"
	"testing"
	"time"
)

func TestEventObserver(t *testing.T) {
	n := observer.EventSubject{
		Observers: sync.Map{},
	}

	tm := time.Now()

	var obs1 = observer.EventObserver{Id: 1, Time: tm}
	var obs2 = observer.EventObserver{Id: 2, Time: tm}
	n.AddListener(&obs1)
	n.AddListener(&obs2)

	go func() {
		select {
		case <-time.After(time.Millisecond * 10):
			n.RemoveListener(&obs1)
		}
	}()

	for x := range observer.Fibonacci(100000) {
		n.Notify(observer.Event{Data: x})
	}
}
