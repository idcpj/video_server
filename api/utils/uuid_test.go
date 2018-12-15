package utils

import (
	"log"
	"testing"
)

func TestNewUuid(t *testing.T) {
	s, e := NewUuid()
	log.Println(s)
	if s == "" || e != nil {
		t.Errorf("%s", e)
	}
}
