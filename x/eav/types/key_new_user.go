package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NewUserKeyPrefix is the prefix to retrieve all NewUser
	NewUserKeyPrefix = "NewUser/value/"
)

// NewUserKey returns the store key to retrieve a NewUser from the index fields
func NewUserKey(
	guid string,
) []byte {
	var key []byte

	guidBytes := []byte(guid)
	key = append(key, guidBytes...)
	key = append(key, []byte("/")...)

	return key
}
