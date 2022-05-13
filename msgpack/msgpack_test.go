package msgpack

import (
	"testing"
)

type data struct {
	Name string
	Age  int
}

var tests = []struct {
	in  data
	out string
}{
	{
		in: data{Name: "foo", Age: 10},
	},
}

func TestEncodingMsgPack(t *testing.T) {
	var msgpack = NewEncodingMsgPack()
	if msgpack == nil {
		t.Fatal("msgpack encoding is nil")
	}

	for _, test := range tests {
		marshal, err := msgpack.Marshal(test.in)
		if err != nil {
			t.Error(err)
		}

		var unmarshal = new(data)
		err = msgpack.Unmarshal(marshal, &unmarshal)
		if err != nil {
			t.Error(err)
		}

		toMap, err := msgpack.UnmarshalToMap(marshal)
		if err != nil {
			t.Error(err)
		}

		if toMap["Name"] != test.in.Name {
			t.Error("Name is not foo")
		}
	}
}
