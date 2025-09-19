package queue

import (
	"math"
	"testing"
)

func TestArrayQueue_EmptyQueueOperationsAndBasicEnqueueDequeue(t *testing.T) {
	q := NewArrayQueue[int]()

	// Test empty queue
	if q.Count() != 0 {
		t.Errorf("Expected count 0, got %d", q.Count())
	}

	// Test dequeue from empty queue
	_, err := q.DeQueue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	// Test enqueue and count
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	if q.Count() != 3 {
		t.Errorf("Expected count 3, got %d", q.Count())
	}
}

func TestArrayQueue_MaintainsFIFOOrderForSequentialOperations(t *testing.T) {
	q := NewArrayQueue[int]()

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

func TestArrayQueue_HandlesCircularBufferWraparoundCorrectly(t *testing.T) {
	q := NewArrayQueue[int]()

	// Fill queue to initial capacity (32)
	for i := 0; i < 32; i++ {
		q.EnQueue(i)
	}

	// Dequeue some elements to create space at front
	for i := 0; i < 10; i++ {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}

	// Add more elements (this should wrap around)
	for i := 32; i < 42; i++ {
		q.EnQueue(i)
	}

	// Verify order is maintained
	expectedValues := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41}

	for _, expected := range expectedValues {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if val != expected {
			t.Errorf("Expected %d, got %d", expected, val)
		}
	}
}

func TestArrayQueue_AutomaticallyResizesWhenCapacityExceeded(t *testing.T) {
	q := NewArrayQueue[int]()

	// Fill beyond initial capacity to trigger resize
	numElements := 100
	for i := 0; i < numElements; i++ {
		q.EnQueue(i)
	}

	if q.Count() != uint(numElements) {
		t.Errorf("Expected count %d, got %d", numElements, q.Count())
	}

	// Verify all elements are in correct order
	for i := 0; i < numElements; i++ {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error at index %d: %v", i, err)
		}
		if val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}
}

func TestArrayQueue_ResizesCorrectlyWhenCircularBufferIsWrapped(t *testing.T) {
	q := NewArrayQueue[int]()

	// Fill to capacity
	for i := 0; i < 32; i++ {
		q.EnQueue(i)
	}

	// Dequeue half to create a wrapped state
	for i := 0; i < 16; i++ {
		q.DeQueue()
	}

	// Fill back up and beyond to trigger resize while wrapped
	for i := 32; i < 50; i++ {
		q.EnQueue(i)
	}

	// Verify order is maintained
	expectedStart := 16
	for i := 0; i < int(q.Count()); i++ {
		val, err := q.DeQueue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		expected := expectedStart + i
		if val != expected {
			t.Errorf("Expected %d, got %d", expected, val)
		}
	}
}

func TestArrayQueue_PanicsWhenExceedingMaximumSize(t *testing.T) {
	q := NewArrayQueue[int]()

	// Get access to the internal structure to manually set count to max
	aq := q.(*arrayQueue[int])
	aq.count = math.MaxUint

	// This should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when enqueueing at max size")
		}
	}()

	q.EnQueue(1)
}

func TestArrayQueue_HandlesInterleavedEnqueueDequeueOperations(t *testing.T) {
	q := NewArrayQueue[string]()

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
}

func TestArrayQueue_PerformsCorrectlyWithLargeNumberOfOperations(t *testing.T) {
	q := NewArrayQueue[int]()

	// Test with a large number of operations
	numOps := 10000

	// Enqueue many items
	for i := 0; i < numOps; i++ {
		q.EnQueue(i)
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
