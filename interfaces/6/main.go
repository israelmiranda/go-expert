package main

import "fmt"

type Observer interface {
	Update(message string)
}

type Subject struct {
	observers []Observer
	message   string
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) SetMessage(message string) {
	s.message = message
	s.notifyAll()
}

func (s *Subject) notifyAll() {
	for _, observer := range s.observers {
		observer.Update(s.message)
	}
}

type ConcreteObserver struct {
	id string
}

func (c ConcreteObserver) Update(message string) {
	fmt.Printf("Observer %s received: %s\n", c.id, message)
}

func main() {
	subject := Subject{}

	observer1 := ConcreteObserver{"A"}
	observer2 := ConcreteObserver{"B"}

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.SetMessage("Hello Observers!")
}
