package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AttributeKeyPrefix is the prefix to retrieve all Attribute
	AttributeKeyPrefix = "Attribute/value/"
)

// AttributeKey returns the store key to retrieve a Attribute from the index fields
func AttributeKey(
	guid string,
) []byte {
	var key []byte

	guidBytes := []byte(guid)
	key = append(key, guidBytes...)
	key = append(key, []byte("/")...)

	return key
}
