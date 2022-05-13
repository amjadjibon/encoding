package toml

import (
	"github.com/pelletier/go-toml"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingTOML is the toml encoding type
type EncodingTOML struct {
}

// Marshal converts the given value to a toml byte slice
func (e *EncodingTOML) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

// Unmarshal converts the given toml byte slice to the given value
func (e *EncodingTOML) Unmarshal(b []byte, v interface{}) error {
	return toml.Unmarshal(b, v)
}

// UnmarshalToMap converts the json string to a map
func (e *EncodingTOML) UnmarshalToMap(b []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	err := e.Unmarshal(b, &m)
	return m, err
}

// NewEncodingTOML creates a new EncodingTOML instance
func NewEncodingTOML() *EncodingTOML {
	return &EncodingTOML{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingTOML, NewEncodingTOML())
}
