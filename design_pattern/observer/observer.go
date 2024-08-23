package oberver

import "fmt"


type IObserver interface {
	Update(msg string)
}

type Observer1 struct {
}

func (o Observer1) Update(msg string) {
	fmt.Println("this is observer1 print msg:"+ msg)
}

type Observer2 struct {}

func(o Observer2) Update(msg string) {
	fmt.Println("this is observer2 print msg:"+ msg)
}

type ISubject interface {
	Attach(observer IObserver)
	Detach(observer IObserver)
	Notify(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Attach(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Detach(observer IObserver) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}



