package random

import mock "github.com/stretchr/testify/mock"

// Mock is an autogenerated mock type for the Interface type
type Mock struct {
	mock.Mock
}

// ExpFloat64 provides a mock function with given fields:
func (m *Mock) ExpFloat64() float64 {
	ret := m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Float32 provides a mock function with given fields:
func (m *Mock) Float32() float32 {
	ret := m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	return r0
}

// Float64 provides a mock function with given fields:
func (m *Mock) Float64() float64 {
	ret := m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Int provides a mock function with given fields:
func (m *Mock) Int() int {
	ret := m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Int31 provides a mock function with given fields:
func (m *Mock) Int31() int32 {
	ret := m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Int31n provides a mock function with given fields: n
func (m *Mock) Int31n(n int32) int32 {
	ret := m.Called(n)

	var r0 int32
	if rf, ok := ret.Get(0).(func(int32) int32); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Int63 provides a mock function with given fields:
func (m *Mock) Int63() int64 {
	ret := m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Int63n provides a mock function with given fields: n
func (m *Mock) Int63n(n int64) int64 {
	ret := m.Called(n)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Intn provides a mock function with given fields: n
func (m *Mock) Intn(n int) int {
	ret := m.Called(n)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NormFloat64 provides a mock function with given fields:
func (m *Mock) NormFloat64() float64 {
	ret := m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Perm provides a mock function with given fields: n
func (m *Mock) Perm(n int) []int {
	ret := m.Called(n)

	var r0 []int
	if rf, ok := ret.Get(0).(func(int) []int); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// Read provides a mock function with given fields: p
func (m *Mock) Read(p []byte) (int, error) {
	ret := m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Seed provides a mock function with given fields: seed
func (m *Mock) Seed(seed int64) {
	m.Called(seed)
}

// Uint32 provides a mock function with given fields:
func (m *Mock) Uint32() uint32 {
	ret := m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}