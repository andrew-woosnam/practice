package queue_impl

import "fmt"

/*
Problem Description

Implement a queue using arrays and pointers. Your implementation should support the following operations:

  - enqueue(val): Add an element to the end of the queue.
  - dequeue(): Remove and return the element from the front of the queue. If the queue is empty, return None.
  - peek(): Return the element at the front of the queue without removing it. If the queue is empty, return None.
  - is_empty(): Return True if the queue is empty, False otherwise.

You must use arrays and pointers to manage the elements of the queue. No other data structures or libraries (e.g., collections.deque, queue.Queue, etc.) are allowed.

Constraints
  - The number of elements in the queue will not exceed 10,000.
*/

type Queue struct {
}

func ArrQueue() (*Queue, error) {
	q := Queue{}
	return &q, nil
}

func (q *Queue) Enqueue(val int) error {
	return fmt.Errorf("not yet implemented")
}

func (q *Queue) Dequeue() (int, error) {
	return -1, fmt.Errorf("not yet implemented")
}

func (q *Queue) Peek() (int, error) {
	return -1, fmt.Errorf("not yet implemented")
}

func (q *Queue) IsEmpty() (bool, error) {
	return false, fmt.Errorf("not yet implemented")
}
