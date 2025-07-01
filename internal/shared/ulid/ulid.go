package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// NewULID generates a new ULID using a monotonic entropy source and the current timestamp.
func NewULID() ulid.ULID {
	// Create a new entropy source for each request
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	// Create a timestamp for ULID (current time)
	ms := ulid.Timestamp(time.Now())
	// Generate a new ULID for the request
	return ulid.MustNew(ms, entropy)
}
