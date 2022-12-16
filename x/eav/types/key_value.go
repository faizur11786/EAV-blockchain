package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValueKeyPrefix is the prefix to retrieve all Value
	ValueKeyPrefix = "Value/value/"
)

// ValueKey returns the store key to retrieve a Value from the index fields
func ValueKey(
	entityId string,
	attributeId string,
) []byte {
	var key []byte

	entityIdBytes := []byte(entityId)
	key = append(key, entityIdBytes...)
	key = append(key, []byte("/")...)

	attributeIdBytes := []byte(attributeId)
	key = append(key, attributeIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
