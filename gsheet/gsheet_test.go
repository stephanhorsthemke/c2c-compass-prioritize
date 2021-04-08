package gsheet

import (
	"testing"
)

var fakeValues [][]interface{} = [][]interface{}{
	{"link1", "DE", "description1", "", "", ""},
	{"link2", "DE", "description2"},
}

func TestParseLinks(t *testing.T) {
	v := parseLinks(fakeValues)

	v0 := Link{0, "link1", "DE", "description1", "", "", ""}
	v1 := Link{0, "link2", "DE", "description2", "", "", ""}

	if v[0] != v0 {
		t.Error("Expected", v0, "got ", v[0].Link)
	}

	if v[1] != v1 {
		t.Error("Expected", v1, "got ", v[1].Link)
	}

}
