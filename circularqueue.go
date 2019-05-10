// Copyright 2019 Dmitry A. Mottl. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.

// Package circularqueue implements a thread-safe circular queue
package circularqueue

import (
	"fmt"
	"strings"
	"sync"
)

// Cicular queue
type Queue struct {
	queue []interface{} // underlying queue slice
	tail  int           // index of the tail
	len   int           // length of the queue
	cap   int           // capacity of the queue
	mutex sync.Mutex
}

// Creates new queue and allocates memory
func NewQueue(capacity int) Queue {
	if capacity <= 0 {
		panic("Capacity should be greater than zero")
	}
	q := Queue{
		tail: -1,
		len:  0,
		cap:  capacity,
	}
	q.queue = make([]interface{}, q.cap, q.cap)
	return q
}

// Returns the length (the number of elements) of the queue
func (q *Queue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.len
}

// Cap returns the capacity (the maximum number of elements) of the queue
func (q *Queue) Cap() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.cap
}

// Vacant returns the number of free slots in the queue
func (q *Queue) Vacant() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.cap - q.len
}

// Push appends the element to the end of the queue
func (q *Queue) Push(element interface{}) (int, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.queue == nil {
		panic("Can't push: queue is uninitialised")
	}
	if q.len == q.cap {
		return 0, fmt.Errorf("Can't push: queue is full (capacity=%d)", q.cap)
	}

	q.len++
	q.tail++
	if q.tail == q.cap {
		q.tail = 0
	}
	q.queue[q.tail] = element
	return q.cap - q.len, nil
}

// PopAt takes and removes the element at the index position from the queue
//  index = 0 to take the first element (the head)
//  index = N to take the Nth element from the head
//  index = -1 to take the last element (the tail)
//  index = -N to take the Nth element from the tail
func (q *Queue) PopAt(index int) (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.queue == nil {
		panic("Can't push: queue is uninitialised")
	}
	if q.len == 0 {
		return nil, fmt.Errorf("Can't pop: queue is empty")
	}

	var realIndex int
	if index < 0 {
		if -index > q.len {
			return nil, fmt.Errorf("Index (%d) exceeds the length of the queue (%d)", index, q.len)
		}
		realIndex = q.tail + index + 1
	} else {
		if index > q.len {
			return nil, fmt.Errorf("Index (%d) exceeds the length of the queue (%d)", index, q.len)
		}
		realIndex = q.tail - q.len + 1 + index
	}

	var element interface{}

	headIndex := q.tail - q.len + 1
	if headIndex < 0 {
		headIndex += q.cap
	}

	if realIndex >= 0 {
		element = q.queue[realIndex]
		if realIndex == q.tail {
			q.queue[q.tail] = nil
			q.tail--
		} else if realIndex == headIndex {
			q.queue[realIndex] = nil
		} else {
			copy(q.queue[realIndex:q.tail], q.queue[realIndex+1:q.tail+1])
			q.queue[q.tail] = nil
			q.tail--
		}
	} else {
		realIndex += q.cap
		element = q.queue[realIndex]
		if realIndex == q.tail {
			q.queue[q.tail] = nil
			q.tail--
		} else if realIndex == headIndex {
			q.queue[realIndex] = nil
		} else {
			copy(q.queue[realIndex:q.cap], q.queue[realIndex+1:q.cap+1]) // copy the end of the slice
			q.queue[q.cap-1] = q.queue[0]
			copy(q.queue[0:q.tail], q.queue[1:q.tail+1]) // copy the start of the slice
			q.queue[q.tail] = nil
			q.tail--
		}
	}

	if q.tail < 0 {
		q.tail += q.cap
	}
	q.len--

	return element, nil
}

// Pops takes and removes the last element from the queue
func (q *Queue) Pop() (interface{}, error) {
	return q.PopAt(-1)
}

// String implements the Stringer interface
func (q Queue) String() string {
	var str strings.Builder

	q.mutex.Lock()
	defer q.mutex.Unlock()

	str.WriteString(fmt.Sprintf("Queue (len=%d, cap=%d) [ ", q.len, q.cap))
	headIndex := q.tail - q.len + 1
	for i := 0; i < q.len; i++ {
		i_ := headIndex + i
		if i_ < 0 {
			i_ += q.cap
		}
		str.WriteString(fmt.Sprintf("%v", q.queue[i_]))
		if i != q.len-1 {
			str.WriteString(", ")
		}
		if (i > 10) && (i < q.len-2) {
			i = q.len - 2
			str.WriteString("..., ")
		}
	}
	str.WriteString(" ]")
	return str.String()
}
