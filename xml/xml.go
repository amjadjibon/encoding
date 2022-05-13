package xml

import (
	"encoding/xml"

	"github.com/clbanning/mxj/v2"

	"github.com/amjadjibon/encoding/constant"
	"github.com/amjadjibon/encoding/registry"
)

// EncodingXML implements the Encoding interface for the XML encoding
type EncodingXML struct {
}

// Marshal converts the given value to YAML byte
func (e *EncodingXML) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Unmarshal converts the given YAML byte to value
func (e *EncodingXML) Unmarshal(b []byte, v interface{}) error {
	return xml.Unmarshal(b, v)
}

// UnmarshalToMap converts the json string to a map
func (e *EncodingXML) UnmarshalToMap(b []byte) (map[string]interface{}, error) {
	var v, err = mxj.NewMapXml(b)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// NewEncodingXML creates a new NewEncodingXML instance
func NewEncodingXML() *EncodingXML {
	return &EncodingXML{}
}

func init() {
	registry.EncodingRegistry().AddEncoding(constant.EncodingXML, NewEncodingXML())
}
