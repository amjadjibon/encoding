package proto

import (
	"github.com/golang/protobuf/proto"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingProto is the proto encoding type
type EncodingProto struct {
}

// Marshal converts the given value to a proto byte slice
func (e *EncodingProto) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

// Unmarshal converts the given proto byte slice to the given value
func (e *EncodingProto) Unmarshal(b []byte, v interface{}) error {
	return proto.Unmarshal(b, v.(proto.Message))
}

// NewEncodingProto creates a new EncodingProto instance
func NewEncodingProto() *EncodingProto {
	return &EncodingProto{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingProto, NewEncodingProto())
}
