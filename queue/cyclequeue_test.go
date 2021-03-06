package queue_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/queue"
)

func TestCycleQueue(t *testing.T) {
	q := queue.NewCycleQueue(10)

	// 入队
	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
		t.Log(q)
	}

	// 出队
	for i := 1; i <= 10; i++ {
		t.Log(q.Dequeue())
		t.Log(q)
	}
}
