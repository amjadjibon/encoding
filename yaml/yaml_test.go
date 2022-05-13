package yaml

import (
	_ "embed"
	"testing"
)

//go:embed test.yaml
var yamlBytes []byte

type Test struct {
	test map[string]interface{}
}

func TestEncodingYAML(t *testing.T) {
	var yaml = NewEncodingYAML()
	var test = make(map[string]interface{})
	err := yaml.Unmarshal(yamlBytes, &test)
	if err != nil {
		t.Error(err)
	}
	if test["name"] != "test" {
		t.Error("name is not equal to test")
	}

	_, err = yaml.Marshal(test)
	if err != nil {
		t.Error(err)
	}

	toMap, err := yaml.UnmarshalToMap(yamlBytes)
	if err != nil {
		t.Error(err)
	}

	if toMap["name"] != "test" {
		t.Error("name is not equal to test")
	}
}
