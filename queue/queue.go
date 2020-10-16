// Package queue implements FIFO data structure using "container/list".
package queue

import (
	"container/list"
	"sync"
)

// Represents FIFO data structure.
type Queue struct {
	// The list
	List *list.List

	// The map
	M map[interface{}]*list.Element

	// Mutex
	Mutex sync.Mutex
}

// Returns initialised queue.
func New() *Queue {
	return &Queue{
		List: list.New(),
		M:    map[interface{}]*list.Element{},
	}
}

// Adds object to the end of a queue.
// All elements of the queue are unique.
// If element already exists in the queue it will be pushed back.
func (q *Queue) Push(object interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if object == nil {
		return
	}

	el, ok := q.M[object]
	if ok {
		q.List.Remove(el)
	}

	q.M[object] = q.List.PushBack(object)
}

// Returns first element of the queue and removes it.
func (q *Queue) Pop() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	el := q.List.Front()
	if el == nil {
		return nil
	}
	q.List.Remove(el)

	return el.Value
}

// Returns the length of the queue
func (q *Queue) Len() int {
	return q.List.Len()
}
