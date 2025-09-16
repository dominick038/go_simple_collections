package stack_test

import (
	"testing"

	"windesheim.dominick038.com/stacksandqueues/stack"
)

func BenchmarkLinkedListStack_Push(b *testing.B) {
	s := stack.NewLinkedListStack[int]()

	for i := 0; b.Loop(); i++ {
		s.Push(i)
	}
}

func BenchmarkLinkedListStack_Pop(b *testing.B) {
	s := stack.NewLinkedListStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		_, err := s.Pop()
		if err != nil {
			b.Fatalf("unexpected error during Pop: %v", err)
		}
	}
}

func BenchmarkLinkedListStack_PushPop(b *testing.B) {
	s := stack.NewLinkedListStack[int]()

	for i := 0; b.Loop(); i++ {
		s.Push(i)
		_, _ = s.Pop()
	}
}
