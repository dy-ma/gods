package queue

type Queue interface {
	Enqueue(interface{}) bool
	Dequeue() (interface{}, bool)
}
