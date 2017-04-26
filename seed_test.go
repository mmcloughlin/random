package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstantSeeder(t *testing.T) {
	seeder := ConstantSeeder(42)
	seed, err := seeder()
	assert.Equal(t, int64(42), seed)
	assert.NoError(t, err)
}
