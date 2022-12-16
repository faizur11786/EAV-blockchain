package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserKeyPrefix is the prefix to retrieve all User
	UserKeyPrefix = "User/value/"
)

// UserKey returns the store key to retrieve a User from the index fields
func UserKey(
	creator string,
	guid string,
) []byte {
	var key []byte

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	guidBytes := []byte(guid)
	key = append(key, guidBytes...)
	key = append(key, []byte("/")...)

	return key
}
