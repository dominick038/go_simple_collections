package stack_test

import (
	"testing"

	"github.com/dominick038/go_simple_collections/stack"
)

func BenchmarkArrayStack_Push(b *testing.B) {
	s := stack.NewArrayStack[int]()

	for i := 0; b.Loop(); i++ {
		s.Push(i)
	}
}

func BenchmarkArrayStack_Pop(b *testing.B) {
	s := stack.NewArrayStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.Pop()
		if err != nil {
			b.Fatalf("unexpected error during Pop: %v", err)
		}
	}
}

func BenchmarkArrayStack_PushPop(b *testing.B) {
	s := stack.NewArrayStack[int]()

	for i := 0; b.Loop(); i++ {
		s.Push(i)
		_, _ = s.Pop()
	}
}
