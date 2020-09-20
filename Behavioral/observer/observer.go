package observer

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type (
	Event struct {
		Data int
	}

	Observer interface {
		NotifyCallback(Event)
	}

	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}

	EventObserver struct {
		Id   int
		Time time.Time
	}

	EventSubject struct {
		Observers sync.Map
	}
)

func (e *EventSubject) AddListener(observer Observer) {
	e.Observers.Store(observer, struct{}{})
}

func (e *EventSubject) RemoveListener(observer Observer) {
	e.Observers.Delete(observer)
	log.Println("==============================")

}

func (e *EventSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		key.(Observer).NotifyCallback(event)
		return true
	})
}

func (e *EventObserver) NotifyCallback(event Event) {
	fmt.Printf("Observer: %d Recieved: %d after %v\n", e.Id, event.Data, time.Since(e.Time))
}

func Fibonacci(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}
