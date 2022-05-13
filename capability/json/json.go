package json

import (
	"encoding/json"

	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/registry"
)

type EncodingJSON struct {
}

func (e *EncodingJSON) Name() string {
	return "encoding_json"
}

func (e *EncodingJSON) Version() string {
	return "1.0.0"
}

func (e *EncodingJSON) Category() string {
	return string(constant.CategoryLibrary)
}

func (e *EncodingJSON) ContractId() string {
	return "encoding:json"
}

func (e *EncodingJSON) New() iface.ICapability {
	return &EncodingJSON{}
}

func (e *EncodingJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (e *EncodingJSON) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func init() {
	registry.GlobalRegistry().AddCapability(&EncodingJSON{})
}
