package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MerchantKeyPrefix is the prefix to retrieve all Merchant
	MerchantKeyPrefix = "Merchant/value/"
)

// MerchantKey returns the store key to retrieve a Merchant from the index fields
func MerchantKey(
	guid string,
	creator string,
) []byte {
	var key []byte

	guidBytes := []byte(guid)
	key = append(key, guidBytes...)
	key = append(key, []byte("/")...)

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
