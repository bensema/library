package example

import (
	"testing"
)

func TestSelect(t *testing.T) {
	err := query()
	if err != nil {
		t.Errorf("Select error:%s", err)
	}
}
