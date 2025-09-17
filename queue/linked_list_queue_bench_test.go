package queue_test

import (
	"testing"

	"github.com/dominick038/go_simple_collections/queue"
)

func BenchmarkLinkedListQueue_EnQueue(b *testing.B) {
	q := queue.NewLinkedListQueue[int]()

	for i := 0; b.Loop(); i++ {
		q.EnQueue(i)
	}
}

func BenchmarkLinkedListQueue_DeQueue(b *testing.B) {
	q := queue.NewLinkedListQueue[int]()
	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
	}

	for i := 0; i < b.N; i++ {
		_, err := q.DeQueue()
		if err != nil {
			b.Fatalf("unexpected error during DeQueue: %v", err)
		}
	}
}

func BenchmarkLinkedListQueue_EnQueueDeQueue(b *testing.B) {
	q := queue.NewLinkedListQueue[int]()

	for i := 0; b.Loop(); i++ {
		q.EnQueue(i)
		_, _ = q.DeQueue()
	}
}
