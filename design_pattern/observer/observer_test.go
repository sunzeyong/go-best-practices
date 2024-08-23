package oberver

import "testing"

func TestSubject(t *testing.T) {
	s := &Subject{}
	s.Attach(Observer1{})
	s.Attach(Observer2{})
	s.Notify("hello")

}
