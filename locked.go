package random

import "sync"

type locked struct {
	base Interface
	lk   sync.Mutex
}

// NewLocked returns a random source that makes locked calls to the given base
// random source, making it safe for use with concurrent goroutines.
func NewLocked(base Interface) Interface {
	return &locked{
		base: base,
	}
}

func (l *locked) Seed(seed int64) {
	l.lk.Lock()
	defer l.lk.Unlock()
	l.base.Seed(seed)
}

func (l *locked) ExpFloat64() float64 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.ExpFloat64()
}

func (l *locked) Float32() float32 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Float32()
}

func (l *locked) Float64() float64 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Float64()
}

func (l *locked) Int() int {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Int()
}

func (l *locked) Int31() int32 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Int31()
}

func (l *locked) Int31n(n int32) int32 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Int31n(n)
}

func (l *locked) Int63() int64 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Int63()
}

func (l *locked) Int63n(n int64) int64 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Int63n(n)
}

func (l *locked) Intn(n int) int {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Intn(n)
}

func (l *locked) NormFloat64() float64 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.NormFloat64()
}

func (l *locked) Perm(n int) []int {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Perm(n)
}

func (l *locked) Uint32() uint32 {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Uint32()
}

func (l *locked) Read(b []byte) (int, error) {
	l.lk.Lock()
	defer l.lk.Unlock()
	return l.base.Read(b)
}
