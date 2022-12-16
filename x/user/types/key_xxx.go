package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// XxxKeyPrefix is the prefix to retrieve all Xxx
	XxxKeyPrefix = "Xxx/value/"
)

// XxxKey returns the store key to retrieve a Xxx from the index fields
func XxxKey(
	address string,
	creator string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
