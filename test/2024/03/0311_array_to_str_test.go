package test

import (
	"strings"
	"testing"
)

func TestSlice2Str(t *testing.T) {
	StringDelimiter := "##qaxbrowser##"
	s := []string{"hello", "world", "gonalng", "dc"}

	str := strings.Join(s, StringDelimiter)
	t.Logf("str = %v", str)
}
