package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MerchantV2KeyPrefix is the prefix to retrieve all MerchantV2
	MerchantV2KeyPrefix = "MerchantV2/value/"
)

// MerchantV2Key returns the store key to retrieve a MerchantV2 from the index fields
func MerchantV2Key(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
