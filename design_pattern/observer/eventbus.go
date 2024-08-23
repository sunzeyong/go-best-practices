package oberver

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	// topic表示一个主题，下面有一系列处理函数，可以将处理函数放入某个主题下面。入参需要相同
	Subscribe(topic string, handler interface{}) error
	// 调用某一个topic下到全部函数
	Public(topic string, args ...interface{})
}

type AsyncEventBus struct {
	Handlers map[string][]reflect.Value
	Lock sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		Handlers: make(map[string][]reflect.Value),
		Lock: sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.Lock.Lock()
	defer bus.Lock.Unlock()

	value := reflect.ValueOf(f)
	if value.Type().Kind() != reflect.Func {
		return fmt.Errorf("this handler is not func")
	}

	handler, ok := bus.Handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, value)
	bus.Handlers[topic] = handler


	return nil
}

func (bus *AsyncEventBus) Public(topic string, args ...interface{}) {
	handler, ok := bus.Handlers[topic]
	if !ok {
		fmt.Println("the topic has empty function")
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for idx := range handler {
		go handler[idx].Call(params)
	}
}


