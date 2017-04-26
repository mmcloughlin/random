package random

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstantSeeder(t *testing.T) {
	seeder := ConstantSeeder(42)
	seed, err := seeder()
	assert.Equal(t, int64(42), seed)
	assert.NoError(t, err)
}

func TestUnixNanoSeeder(t *testing.T) {
	seed, err := UnixNanoSeeder.Seed()
	require.NoError(t, err)
	now := time.Now().UnixNano()
	assert.True(t, seed <= now)
}

func TestReaderSeeder(t *testing.T) {
	buf := bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	seeder := NewReaderSeeder(buf)

	seed, err := seeder.Seed()
	require.NoError(t, err)
	assert.Equal(t, int64(0x0102030405060708), seed)

	// 2 bytes left, so expect short read
	_, err = seeder.Seed()
	require.Error(t, err)
}

func TestNew(t *testing.T) {
	r1 := New()
	r2 := New()
	assert.NotEqual(t, r1.Int(), r2.Int())
}

func TestNewWithSeed(t *testing.T) {
	r := NewWithSeed(42)
	x := r.Intn(10000)
	assert.Equal(t, 2305, x)
}

func TestNewFromSeederError(t *testing.T) {
	seeder := &MockSeeder{}
	seeder.On("Seed").Return(int64(0), assert.AnError)
	_, err := NewFromSeeder(seeder)
	assert.Error(t, err)
}

func TestMust(t *testing.T) {
	expect := &Mock{}
	r := Must(expect, nil)
	assert.Equal(t, expect, r)
}

func TestMustPanics(t *testing.T) {
	assert.Panics(t, func() {
		Must(nil, assert.AnError)
	})
}
