package queue

import (
	"testing"
)

func TestLinkedListQueue_EmptyQueueOperationsAndBasicEnqueueDequeue(t *testing.T) {
	q := NewLinkedListQueue[int]()

	// Test empty queue
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}

	// Test dequeue from empty queue
	_, err := q.DeQueue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	// Test error type
	if _, ok := err.(*QueueEmptyError); !ok {
		t.Errorf("Expected QueueEmptyError, got %T", err)
	}

	// Test enqueue and count
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	if q.Count() != 3 {
		t.Errorf("Expected count 3, got %d", q.Count())
	}
}

func TestLinkedListQueue_MaintainsFIFOOrderForSequentialOperations(t *testing.T) {
	q := NewLinkedListQueue[int]()

	// Enqueue values 1-10
	for i := 1; i <= 10; i++ {
		q.EnQueue(i)
	}

	// Dequeue and verify FIFO order
	for i := 1; i <= 10; i++ {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}

	// Queue should be empty now
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}
}

func TestLinkedListQueue_SingleElementOperations(t *testing.T) {
	q := NewLinkedListQueue[string]()

	// Add single element
	q.EnQueue("test")
	if q.Count() != 1 {
		t.Errorf("Expected count 1, got %d", q.Count())
	}

	// Remove single element
	val, err := q.DeQueue()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if val != "test" {
		t.Errorf("Expected 'test', got '%s'", val)
	}

	// Should be empty
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}

	// Should error on next dequeue
	_, err = q.DeQueue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}
}

func TestLinkedListQueue_HandlesInterleavedEnqueueDequeueOperations(t *testing.T) {
	q := NewLinkedListQueue[string]()

	// Test with strings to ensure generics work
	testData := []string{"hello", "world", "test", "queue"}

	// Enqueue all
	for _, str := range testData {
		q.EnQueue(str)
	}

	// Dequeue two
	val1, err := q.DeQueue()
	if err != nil || val1 != "hello" {
		t.Errorf("Expected 'hello', got '%s'", val1)
	}

	val2, err := q.DeQueue()
	if err != nil || val2 != "world" {
		t.Errorf("Expected 'world', got '%s'", val2)
	}

	// Add more
	q.EnQueue("new")
	q.EnQueue("items")

	// Dequeue remaining in order
	expected := []string{"test", "queue", "new", "items"}
	for _, exp := range expected {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if val != exp {
			t.Errorf("Expected '%s', got '%s'", exp, val)
		}
	}

	// Should be empty
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}
}

func TestLinkedListQueue_AlternatingEnqueueDequeue(t *testing.T) {
	q := NewLinkedListQueue[int]()

	// Test alternating operations with predictable pattern
	// Add 1, 2, remove 1, add 3, 4, remove 2, etc.
	expected := []int{}

	for i := 0; i < 10; i++ {
		// Add two elements
		val1 := i*2 + 1
		val2 := i*2 + 2
		q.EnQueue(val1)
		q.EnQueue(val2)
		expected = append(expected, val1, val2)

		if i > 0 { // Don't dequeue on first iteration
			// Remove one element
			removed, err := q.DeQueue()
			if err != nil {
				t.Fatalf("Unexpected error at iteration %d: %v", i, err)
			}
			if removed != expected[0] {
				t.Errorf("Expected to remove %d, got %d", expected[0], removed)
			}
			expected = expected[1:] // Remove first element from expected
		}
	}

	// Drain all remaining elements
	for len(expected) > 0 {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error draining: %v", err)
		}
		if val != expected[0] {
			t.Errorf("Expected %d, got %d", expected[0], val)
		}
		expected = expected[1:]
	}

	// Should be empty
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}
}

func TestLinkedListQueue_PerformsCorrectlyWithLargeNumberOfOperations(t *testing.T) {
	q := NewLinkedListQueue[int]()

	// Test with a large number of operations
	numOps := 10000

	// Enqueue many items
	for i := 0; i < numOps; i++ {
		q.EnQueue(i)
	}

	if q.Count() != uint(numOps) {
		t.Errorf("Expected count %d, got %d", numOps, q.Count())
	}

	// Dequeue half
	for i := 0; i < numOps/2; i++ {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}

	// Enqueue more
	for i := numOps; i < numOps+numOps/2; i++ {
		q.EnQueue(i)
	}

	// Dequeue all remaining
	expectedStart := numOps / 2
	for i := 0; i < numOps; i++ {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		expected := expectedStart + i
		if val != expected {
			t.Errorf("Expected %d, got %d", expected, val)
		}
	}

	// Should be empty
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}
}

func TestLinkedListQueue_WorksWithDifferentPrimitiveTypes(t *testing.T) {
	// Test with int
	intQ := NewLinkedListQueue[int]()
	intQ.EnQueue(42)
	val, _ := intQ.DeQueue()
	if val != 42 {
		t.Errorf("Int queue: expected 42, got %d", val)
	}

	// Test with float64
	floatQ := NewLinkedListQueue[float64]()
	floatQ.EnQueue(3.14)
	fVal, _ := floatQ.DeQueue()
	if fVal != 3.14 {
		t.Errorf("Float queue: expected 3.14, got %f", fVal)
	}

	// Test with bool
	boolQ := NewLinkedListQueue[bool]()
	boolQ.EnQueue(true)
	boolQ.EnQueue(false)
	b1, _ := boolQ.DeQueue()
	b2, _ := boolQ.DeQueue()
	if b1 != true || b2 != false {
		t.Errorf("Bool queue: expected true,false got %t,%t", b1, b2)
	}

	// Test with string
	strQ := NewLinkedListQueue[string]()
	strQ.EnQueue("golang")
	sVal, _ := strQ.DeQueue()
	if sVal != "golang" {
		t.Errorf("String queue: expected 'golang', got '%s'", sVal)
	}
}

func TestLinkedListQueue_CountAccuracyThroughoutOperations(t *testing.T) {
	q := NewLinkedListQueue[int]()

	// Start empty
	if q.Count() != 0 {
		t.Errorf("Initial count should be 0, got %d", q.Count())
	}

	// Add elements and check count increases
	for i := 1; i <= 5; i++ {
		q.EnQueue(i)
		if q.Count() != uint(i) {
			t.Errorf("After adding %d elements, expected count %d, got %d", i, i, q.Count())
		}
	}

	// Remove elements and check count decreases
	for i := 4; i >= 0; i-- {
		q.DeQueue()
		if q.Count() != uint(i) {
			t.Errorf("After removing element, expected count %d, got %d", i, q.Count())
		}
	}

	// Should be empty again
	if q.Count() != 0 {
		t.Errorf("Final count should be 0, got %d", q.Count())
	}
}

func TestLinkedListQueue_PreservesOrderWithRepeatedValues(t *testing.T) {
	q := NewLinkedListQueue[int]()

	// Add repeated values
	values := []int{1, 1, 2, 2, 3, 3, 1, 2, 3}
	for _, val := range values {
		q.EnQueue(val)
	}

	// Dequeue and verify order is preserved
	for i, expected := range values {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error at position %d: %v", i, err)
		}
		if val != expected {
			t.Errorf("Position %d: expected %d, got %d", i, expected, val)
		}
	}
}
