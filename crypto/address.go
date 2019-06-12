package crypto

import (
	"encoding/hex"
)

const (
	// AddressLength is the size of an account address.
	AddressLength = 20
)

// Address represents the 20 byte address of an account.
type Address [AddressLength]byte

// BytesToAddress returns Address with value b.
// If b is larger than len(h), b will be cropped from the left.
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

// SetBytes sets the address to the value of b.
// If b is larger than len(a) it will panic.
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

// Bytes gets the byte representation of the underlying address.
func (a Address) Bytes() []byte { return a[:] }

// String returns the string encoding of the address.
func (a Address) String() string {
	return "0x" + hex.EncodeToString(a.Bytes())
}

// ZeroAddress represents the "zero address" (account that no one owns).
func ZeroAddress() Address {
	return Address{}
}