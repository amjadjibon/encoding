package xml

import (
	_ "embed"
	"testing"
)

//go:embed test.xml
var xmlBytes []byte

type Plant struct {
	Id     int      `xml:"id,attr"`
	Name   string   `xml:"name"`
	Origin []string `xml:"origin"`
}

func TestEncodingXML(t *testing.T) {
	var xml = NewEncodingXML()
	if xml == nil {
		t.Error("NewEncodingXML() returned nil")
	}
	var test = &Plant{Id: 1, Name: "Plant 1", Origin: []string{"China", "Japan"}}

	marshal, err := xml.Marshal(test)
	if err != nil {
		t.Error(err)
	}

	var test2 = &Plant{}
	err = xml.Unmarshal(marshal, test2)
	if err != nil {
		t.Error(err)
	}

	if test2.Id != test.Id {
		t.Errorf("test2.Id = %d; want %d", test2.Id, test.Id)
	}
}

func TestEncodingXML_UnmarshalToMap(t *testing.T) {
	var xml = NewEncodingXML()
	if xml == nil {
		t.Error("NewEncodingXML() returned nil")
	}

	var _, err = xml.UnmarshalToMap(xmlBytes)
	if err != nil {
		t.Error(err)
	}

	_, err = xml.UnmarshalToMap(nil)
	if err == nil {
		t.Error("UnmarshalToMap(nil) should return error")
	}
}
