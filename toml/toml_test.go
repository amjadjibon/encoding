package toml

import (
	_ "embed"
	"testing"
)

//go:embed test.toml
var tomlBytes []byte

func TestEncodingTOML(t *testing.T) {
	var toml = NewEncodingTOML()
	if toml == nil {
		t.Error("NewEncodingTOML() returned nil")
	}

	var m = make(map[string]interface{})

	err := toml.Unmarshal(tomlBytes, &m)
	if err != nil {
		t.Error("Unmarshal() returned an error:", err)
	}

	if m["title"] != "TOML Example" {
		t.Error("Unmarshal() did not set the correct value for 'name'")
	}

	if m["owner"].(map[string]interface{})["name"] != "Tom Preston-Werner" {
		t.Error("Unmarshal() did not set the correct value for 'owner.name'")
	}

	_, err = toml.Marshal(m)
	if err != nil {
		t.Error(err)
	}

	toMap, err := toml.UnmarshalToMap(tomlBytes)
	if err != nil {
		t.Error(err)
	}

	if toMap["title"] != "TOML Example" {
		t.Error("UnmarshalToMap() did not set the correct value for 'name'")
	}
}
