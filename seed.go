package random

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"io"
	"math/rand"
	"time"
)

// Seeder is a method of generating a seed.
type Seeder interface {
	Seed() (int64, error)
}

// SeedFunc is an adapter to allow ordinary functions to be used as Seeder
// implementations.
type SeedFunc func() (int64, error)

// Confirm SeedFunc satisfies Seeder interface.
var _ Seeder = new(SeedFunc)

// Seed calls f.
func (f SeedFunc) Seed() (int64, error) {
	return f()
}

// ConstantSeeder builds a SeedFunc that always returns seed.
func ConstantSeeder(seed int64) SeedFunc {
	return func() (int64, error) { return seed, nil }
}

// UnixNanoSeeder generates seeds with the current nanosecond epoch.
var UnixNanoSeeder = SeedFunc(func() (int64, error) {
	return time.Now().UnixNano(), nil
})

// ReaderSeeder reads seeds from an underlying io.Reader.
type ReaderSeeder struct {
	reader io.Reader
}

// Confirm ReaderSeeder is a Seeder.
var _ Seeder = new(ReaderSeeder)

// NewReaderSeeder builds a ReaderSeeder reading from r.
func NewReaderSeeder(r io.Reader) Seeder {
	return &ReaderSeeder{
		reader: r,
	}
}

// Seed reads 8 bytes from the underlying reader and generates a seed from it.
func (r *ReaderSeeder) Seed() (int64, error) {
	buf := [8]byte{}
	_, err := io.ReadFull(r.reader, buf[:])
	if err != nil {
		return 0, err
	}
	seed := binary.BigEndian.Uint64(buf[:])
	return int64(seed), nil
}

// CryptoSeeder generates seeds using crypto/rand.
var CryptoSeeder = NewReaderSeeder(cryptorand.Reader)

// New builds a new random generator (seeded with current nanosecond unix
// epoch).
func New() Interface {
	return Must(NewFromSeeder(UnixNanoSeeder))
}

// NewWithSeed builds a random generator with the given seed.
func NewWithSeed(seed int64) Interface {
	return Must(NewFromSeeder(ConstantSeeder(seed)))
}

// NewFromSeeder builds a new random generator with a seed generated with s.
func NewFromSeeder(s Seeder) (Interface, error) {
	seed, err := s.Seed()
	if err != nil {
		return nil, err
	}
	return rand.New(rand.NewSource(seed)), nil
}

// Must panics if err is non-nil and returns r otherwise.
func Must(r Interface, err error) Interface {
	if err != nil {
		panic(err)
	}
	return r
}
