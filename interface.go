package random

import (
	"io"
	"math/rand"
)

// Interface is an interface for random generators.
type Interface interface {
	Seed(seed int64)
	ExpFloat64() float64
	Float32() float32
	Float64() float64
	Int() int
	Int31() int32
	Int31n(n int32) int32
	Int63() int64
	Int63n(n int64) int64
	Intn(n int) int
	NormFloat64() float64
	Perm(n int) []int
	Uint32() uint32
	io.Reader
}

// Confirm the standard library Rand object conforms to our Interface.
var _ Interface = new(rand.Rand)
