package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingMsgPack is the msgpack encoding registry used by the vmihailenco/msgpack package
type EncodingMsgPack struct{}

// Marshal converts the given value to a msgpack byte slice
func (e *EncodingMsgPack) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal converts the given msgpack byte slice to the given value
func (e *EncodingMsgPack) Unmarshal(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b, v)
}

// UnmarshalToMap converts the json string to a map
func (e *EncodingMsgPack) UnmarshalToMap(b []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	err := e.Unmarshal(b, &m)
	return m, err
}

// NewEncodingMsgPack returns a new EncodingMsgPack
func NewEncodingMsgPack() *EncodingMsgPack {
	return &EncodingMsgPack{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingMsgPack, NewEncodingMsgPack())
}
