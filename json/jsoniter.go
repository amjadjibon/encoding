package json

import (
	jsonIter "github.com/json-iterator/go"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingJSONIter is a json encoding using json-iter
type EncodingJSONIter struct {
}

// Marshal returns the EncodingJSON encoding of v
func (e EncodingJSONIter) Marshal(v interface{}) ([]byte, error) {
	return jsonIter.Marshal(v)
}

// Unmarshal parses the EncodingJSON-encoded data and stores the result
func (e EncodingJSONIter) Unmarshal(b []byte, v interface{}) error {
	return jsonIter.Unmarshal(b, v)
}

// UnmarshalToMap converts the json string to a map
func (e *EncodingJSONIter) UnmarshalToMap(b []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	err := jsonIter.Unmarshal(b, &m)
	return m, err
}

// NewJEncodingJSONIter returns a new EncodingJSONIter
func NewJEncodingJSONIter() *EncodingJSONIter {
	return &EncodingJSONIter{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingJSONIter, NewJEncodingJSONIter())
}
