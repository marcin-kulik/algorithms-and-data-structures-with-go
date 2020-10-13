package queue

import "container/list"

type FifoImpl struct {
	List *list.List
}

func NewFifoImpl() FifoImpl {
	return FifoImpl{
		List: list.New(),
	}
}

func (q FifoImpl) Push(object interface{}) {
	q.List.PushBack(object)
}

func (q FifoImpl) Pop() interface{} {

	el := q.List.Front()
	if el == nil {
		return nil
	}
	q.List.Remove(el)

	return el.Value
}
