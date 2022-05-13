package gob

import (
	"testing"
)

func TestEncodingGOB(t *testing.T) {
	var gob = NewEncodingGOB()
	if gob == nil {
		t.Error("NewEncodingGOB() returned nil")
	}

	var data = struct {
		A int
		B string
	}{
		A: 1,
		B: "2",
	}

	var encoded []byte
	var err error

	if encoded, err = gob.Marshal(data); err != nil {
		t.Errorf("Marshal() returned error: %s", err)
	}

	err = gob.Unmarshal(encoded, &data)
	if err != nil {
		t.Error("Unmarshal() returned error:", err)
	}

	if encoded, err = gob.Marshal(nil); err == nil {
		t.Errorf("Marshal() should return error")
	}

	if err = gob.Unmarshal(nil, &data); err == nil {
		t.Errorf("Unmarshal() should return error")
	}
}
