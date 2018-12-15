package utils

import (
	"testing"
)

func TestNewUuid(t *testing.T) {
	s, e := NewUuid()
	if s == "" || e != nil {
		t.Errorf("%s", e)
	}
}
