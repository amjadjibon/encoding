package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingGOB is the GOB encoding type
type EncodingGOB struct {
}

// Marshal converts the input object to a byte array
func (e *EncodingGOB) Marshal(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	var encoder = gob.NewEncoder(&buffer)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Unmarshal converts the input byte array to an object
func (e *EncodingGOB) Unmarshal(data []byte, v interface{}) error {
	var reader = bytes.NewReader(data)
	var decoder = gob.NewDecoder(reader)
	return decoder.Decode(v)
}

// NewEncodingGOB returns a new EncodingGOB
func NewEncodingGOB() *EncodingGOB {
	return &EncodingGOB{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingGOB, NewEncodingGOB())
}
