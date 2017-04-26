package random

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SumSerial(r Interface, n int) int64 {
	sum := int64(0)
	for i := 0; i < n; i++ {
		sum += r.Int63()
	}
	return sum
}

func SumParallel(r Interface, m, n int) int64 {
	sum := int64(0)
	var wg sync.WaitGroup
	wg.Add(m)
	for i := 0; i < m; i++ {
		go func() {
			sub := SumSerial(r, n)
			atomic.AddInt64(&sum, sub)
			wg.Done()
		}()
	}

	wg.Wait()
	return sum
}

// TestLocked would fail without wrapping the random generator passed to
// SumParallel in NewLocked.
func TestLocked(t *testing.T) {
	seed := int64(42)
	m := 1000
	n := 1000
	N := m * n

	r := NewLocked(NewWithSeed(seed)) // NewLocked critical here
	got := SumParallel(r, m, n)
	t.Logf("parallel sum seed=%d m=%d n=%d result=%d", seed, m, n, got)

	r = NewWithSeed(seed)
	expect := SumSerial(r, N)
	t.Logf("serial sum seed=%d N=%d result=%d", seed, N, expect)

	assert.Equal(t, expect, got)
}
