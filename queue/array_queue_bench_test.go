package queue_test

import (
	"testing"

	"github.com/dominick038/go_simple_collections/queue"
)

func BenchmarkArrayQueue_EnQueue(b *testing.B) {
	q := queue.NewArrayQueue[int]()

	for i := 0; b.Loop(); i++ {
		q.EnQueue(i)
	}
}

func BenchmarkArrayQueue_DeQueue(b *testing.B) {
	q := queue.NewArrayQueue[int]()
	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := q.DeQueue()
		if err != nil {
			b.Fatalf("unexpected error during DeQueue: %v", err)
		}
	}
}

func BenchmarkArrayQueue_EnQueueDeQueue(b *testing.B) {
	q := queue.NewArrayQueue[int]()

	for i := 0; b.Loop(); i++ {
		q.EnQueue(i)
		_, _ = q.DeQueue()
	}
}
