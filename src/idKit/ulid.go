package idKit

import (
	"github.com/oklog/ulid/v2"
	"io"
)

// NewULID
/*
PS:
If you just want to generate a ULID and don't (yet) care about details like performance, cryptographic security, etc., use the ulid.Make helper function.
This function calls time.Now to get a timestamp, and uses a source of entropy which is process-global, pseudo-random, and monotonic.
*/
func NewULID() string {
	return ulid.Make().String()
}

func NewCustomizedULID(ms uint64, entropy io.Reader) (string, error) {
	//entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	//ms := ulid.Timestamp(time.Now())

	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
