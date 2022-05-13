package json

import (
	"encoding/json"

	"github.com/amjadjibon/encoding/registry"
)

type EncodingJSON struct {
}

func (e *EncodingJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (e *EncodingJSON) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func init() {
	registry.EncodingRegistry().AddEncoding("json", &EncodingJSON{})
}
