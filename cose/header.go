// (c) 2022-2022, LDC Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cose

import "github.com/ldclabs/cose/key"

// Headers represents a COSE Generic_Headers structure.
type Headers key.IntMap

// Has returns true if the Headers has the given parameter.
func (h Headers) Has(p int) bool {
	return key.IntMap(h).Has(p)
}

// GetBool returns the value of the given parameter as a bool, or a error.
func (h Headers) GetBool(p int) (bool, error) {
	return key.IntMap(h).GetBool(p)
}

// GetInt returns the value of the given parameter as a int, or a error.
func (h Headers) GetInt(p int) (int, error) {
	return key.IntMap(h).GetInt(p)
}

// GetInt64 returns the value of the given parameter as a int64, or a error.
func (h Headers) GetInt64(p int) (int64, error) {
	return key.IntMap(h).GetInt64(p)
}

// GetUint64 returns the value of the given parameter as a uint64, or a error.
func (h Headers) GetUint64(p int) (uint64, error) {
	return key.IntMap(h).GetUint64(p)
}

// GetBytes returns the value of the given parameter as a slice of bytes, or a error.
func (h Headers) GetBytes(p int) ([]byte, error) {
	return key.IntMap(h).GetBytes(p)
}

// GetString returns the value of the given parameter as a string, or a error.
func (h Headers) GetString(p int) (string, error) {
	return key.IntMap(h).GetString(p)
}

// Bytesify returns a CBOR-encoded byte slice.
// It returns nil if MarshalCBOR failed.
func (h Headers) Bytesify() []byte {
	b, _ := key.MarshalCBOR(h)
	return b
}
