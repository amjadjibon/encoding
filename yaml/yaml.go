package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingYAML implements the Encoding interface for the YAML encoding
type EncodingYAML struct {
}

// Marshal converts the given value to YAML byte
func (e *EncodingYAML) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// UnmarshalToMap converts the json string to a map
func (e *EncodingYAML) UnmarshalToMap(b []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	err := e.Unmarshal(b, &m)
	return m, err
}

// Unmarshal converts the given YAML byte to value
func (e *EncodingYAML) Unmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

// NewEncodingYAML creates a new EncodingYAML instance.
func NewEncodingYAML() *EncodingYAML {
	return &EncodingYAML{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingYAML, NewEncodingYAML())
}
