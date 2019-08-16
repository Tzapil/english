package common

import (
	"encoding/hex"
	"errors"
)

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

type ObjectID [12]byte

var NilObjectID ObjectID

// FromHex creates a new ObjectID from a hex string. It returns an error if the hex string is not a
// valid ObjectID.
func FromHex(s string) (ObjectID, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return NilObjectID, err
	}

	if len(b) != 12 {
		return NilObjectID, errors.New("the provided hex string is not a valid ObjectID")
	}

	var oid [12]byte
	copy(oid[:], b[:])

	return oid, nil
}
