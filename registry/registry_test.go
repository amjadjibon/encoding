package registry

import (
	"testing"

	"github.com/amjadjibon/encoding/iface"
)

type TestEncoding struct {
}

func (t TestEncoding) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (t TestEncoding) Unmarshal(b []byte, v interface{}) error {
	return nil
}

func TestEncodingRegistry(t *testing.T) {
	registry := EncodingRegistry()
	if registry == nil {
		t.Fatal("EncodingRegistry() returned nil")
	}

	registry.AddEncoding("test", &TestEncoding{})

	encoding := registry.GetEncoding("test")

	if encoding == nil {
		t.Fatal("GetEncoding() returned nil")
	}

	if _, ok := encoding.(iface.IEncoding); !ok {
		t.Fatal("GetEncoding() returned an invalid interface")
	}
}
