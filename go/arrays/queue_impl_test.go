package queue_impl

import "testing"

// Helper function to report test failure
func fail(t *testing.T, format string, args ...interface{}) {
	t.Fatalf(format, args...)
}

// AssertEqual checks if the value matches the expected value
func AssertEqual(expected, actual interface{}, t *testing.T) {
	if actual != expected {
		fail(t, "expected %v, got %v", expected, actual)
	}
}

// Helper function to initialize the queue
func initArrQueue(t *testing.T) *Queue {
	q, err := ArrQueue()
	if err != nil {
		t.Fatalf("Failed to initialize ArrQueue: %v", err)
	}
	return q
}

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := initArrQueue(t)

	AssertEqual(nil, q.Enqueue(1), t)
	AssertEqual(nil, q.Enqueue(2), t)
	AssertEqual(nil, q.Enqueue(3), t)

	val, err := q.Dequeue()
	AssertEqual(nil, err, t)
	AssertEqual(1, val, t)

	val, err = q.Dequeue()
	AssertEqual(nil, err, t)
	AssertEqual(2, val, t)

	val, err = q.Dequeue()
	AssertEqual(nil, err, t)
	AssertEqual(3, val, t)

	_, err = q.Dequeue()
	AssertEqual(err != nil, true, t)
}

func TestQueue_Peek(t *testing.T) {
	q := initArrQueue(t)

	AssertEqual(nil, q.Enqueue(1), t)
	AssertEqual(nil, q.Enqueue(2), t)
	AssertEqual(nil, q.Enqueue(3), t)

	val, err := q.Peek()
	AssertEqual(nil, err, t)
	AssertEqual(1, val, t)

	q.Dequeue()
	val, err = q.Peek()
	AssertEqual(nil, err, t)
	AssertEqual(2, val, t)

	q.Dequeue()
	val, err = q.Peek()
	AssertEqual(nil, err, t)
	AssertEqual(3, val, t)

	q.Dequeue()
	_, err = q.Peek()
	AssertEqual(err != nil, true, t)
}

func TestQueue_IsEmpty(t *testing.T) {
	q := initArrQueue(t)

	empty, err := q.IsEmpty()
	AssertEqual(nil, err, t)
	AssertEqual(true, empty, t)

	AssertEqual(nil, q.Enqueue(1), t)
	empty, err = q.IsEmpty()
	AssertEqual(nil, err, t)
	AssertEqual(false, empty, t)

	q.Dequeue()
	empty, err = q.IsEmpty()
	AssertEqual(nil, err, t)
	AssertEqual(true, empty, t)
}

func TestQueue_LargeNumberOfOperations(t *testing.T) {
	q := initArrQueue(t)

	for i := 0; i < 10000; i++ {
		AssertEqual(nil, q.Enqueue(i), t)
	}

	for i := 0; i < 10000; i++ {
		val, err := q.Dequeue()
		AssertEqual(nil, err, t)
		AssertEqual(i, val, t)
	}

	empty, err := q.IsEmpty()
	AssertEqual(nil, err, t)
	AssertEqual(true, empty, t)
}

func TestQueue_EmptyQueueOperations(t *testing.T) {
	q := initArrQueue(t)

	_, err := q.Dequeue()
	AssertEqual(err != nil, true, t)

	_, err = q.Peek()
	AssertEqual(err != nil, true, t)
}
