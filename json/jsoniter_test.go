package json

import (
	"testing"
)

func TestEncodingJSONIter(t *testing.T) {
	var json = NewJEncodingJSONIter()
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

func TestEncodingJSONIter_UnmarshalToMap(t *testing.T) {
	var json = NewJEncodingJSONIter()
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
