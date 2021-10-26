package queue

import "sync"

type Queue struct {
	// v holds the queue values as an interface array
	// so, we can add any kind of item into the list.
	v []interface{}

	// count holds the queue length inside the Queue struct
	count int

	// mutex satisfy the synchronization for queue usage in goroutines
	mu sync.RWMutex
}

// NewQueue inits the queue
func NewQueue() *Queue {
	return &Queue{
		v: make([]interface{}, 0),
	}
}

// Push adds an item to the list(queue)
// Then increses count of the Queue
func (q *Queue) Push(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.v = append(q.v, v)
	q.count++
}

// Pop gets the first added item from the list
// Also descreases the count of queue
// if count isn't 0(zero) already
func (q *Queue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isEmpty() {
		return nil
	}

	v := q.v[0]
	// if list have only 1 item
	// then it will panic if we try to get non-existing array elements
	// and error will be : slice bounds out of range
	if q.count == 1 {
		q.v = q.v[:0] // return zero-valued array
	} else {
		q.v = q.v[1:]
	}
	q.count--

	return v
}

// Count return the length of the queue
func (q *Queue) Count() int {
	return q.count
}

// isEmpty checks the queue if it has any item in it.
func (q *Queue) isEmpty() bool {
	return q.count == 0
}
