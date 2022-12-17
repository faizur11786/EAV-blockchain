package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MerchantNewKeyPrefix is the prefix to retrieve all MerchantNew
	MerchantNewKeyPrefix = "MerchantNew/value/"
)

// MerchantNewKey returns the store key to retrieve a MerchantNew from the index fields
func MerchantNewKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
