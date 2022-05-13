package json

import (
	"encoding/json"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingJSON is the json encoding type
type EncodingJSON struct {
}

// Marshal converts the given value to a json byte slice
func (e *EncodingJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal converts the given json byte slice to the given value
func (e *EncodingJSON) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

// UnmarshalToMap converts the json string to a map
func (e *EncodingJSON) UnmarshalToMap(b []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	return m, err
}

// NewEncodingJSON creates a new EncodingJSON instance
func NewEncodingJSON() *EncodingJSON {
	return &EncodingJSON{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingJSON, &EncodingJSON{})
}
