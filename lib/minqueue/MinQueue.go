package minqueue

// Comparable defines an interface for accepted types in the Min Queue.
type Comparable interface {
	LessThan(other Comparable) bool
}

type MinQueue interface {
	IsEmpty() bool
	Insert(val Comparable) bool
	Pull() (Comparable, bool)
}
