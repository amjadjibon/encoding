package proto

import (
	"testing"
)

func TestEncodingProto(t *testing.T) {
	var kv = &KeyValueStore{
		Key:   "key",
		Value: "value",
	}

	var proto = NewEncodingProto()
	if proto == nil {
		t.Error("NewEncodingProto() returned nil")
	}

	marshal, err := proto.Marshal(kv)
	if err != nil {
		t.Error("proto.Marshal() returned an error")
	}

	if marshal == nil {
		t.Error("proto.Marshal() returned nil")
	}

	var unmarshal = &KeyValueStore{}
	err = proto.Unmarshal(marshal, unmarshal)
	if err != nil {
		t.Error("proto.Unmarshal() returned an error")
	}

	if unmarshal.GetKey() != kv.Key {
		t.Error("proto.Unmarshal() returned a different key")
	}

	if unmarshal.GetValue() != kv.Value {
		t.Error("proto.Unmarshal() returned a different value")
	}
}
