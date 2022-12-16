package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EntityTypeKeyPrefix is the prefix to retrieve all EntityType
	EntityTypeKeyPrefix = "EntityType/value/"
)

// EntityTypeKey returns the store key to retrieve a EntityType from the index fields
func EntityTypeKey(
	guid string,
) []byte {
	var key []byte

	guidBytes := []byte(guid)
	key = append(key, guidBytes...)
	key = append(key, []byte("/")...)

	return key
}
