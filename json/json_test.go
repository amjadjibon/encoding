package json

import (
	_ "embed"
	"testing"
)

//go:embed test.json
var jsonBytes []byte

var tests = []struct {
	in  interface{}
	out string
}{
	{
		in: struct {
			Name string
			Age  int
		}{
			Name: "John",
			Age:  30,
		},

		out: `{"Name":"John","Age":30}`,
	},
}

func TestEncodingJSON(t *testing.T) {
	var json = NewEncodingJSON()
	if json == nil {
		t.Fatal("json is nil")
	}

	for _, test := range tests {
		marshal, err := json.Marshal(test.in)
		if err != nil {
			t.Error(err)
		}
		if string(marshal) != test.out {
			t.Errorf("expected %s, got %s", test.out, marshal)
		}

		unmarshal := &test.in
		err = json.Unmarshal([]byte(test.out), unmarshal)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestEncodingJSON_UnmarshalToMap(t *testing.T) {
	var json = NewEncodingJSON()
	if json == nil {
		t.Fatal("json is nil")
	}

	toMap, err := json.UnmarshalToMap(jsonBytes)
	if err != nil {
		t.Error(err)
	}

	if len(toMap) != 1 {
		t.Errorf("expected 1, got %d", len(toMap))
	}

	if toMap["student"].([]interface{})[0].(map[string]interface{})["name"] != "Tom" {
		t.Errorf("expected test, got %s", toMap["test"])
	}
}
