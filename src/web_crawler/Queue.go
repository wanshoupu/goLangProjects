package main

import "sync"

func NewQueue(capacity int) *Queue {
	return &Queue{
		Items: make([]interface{}, 0, capacity),
	}
}

type Queue struct {
	sync.Mutex
	Items []interface{}
}

func (q *Queue) Len() int {
	return len(q.Items)
}

func (q *Queue) Push(item interface{}) {
	q.Lock()
	defer q.Unlock()
	q.Items = append(q.Items, item)
}

func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
