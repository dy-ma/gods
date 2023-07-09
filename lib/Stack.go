package lib

type Stack interface {
	Push(interface{}) bool
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)
}
