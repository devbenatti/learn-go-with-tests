package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	greet(&buffer, "Chris")
	result := buffer.String()
	expect := "Hello, Chris"

	if result != expect {
		t.Errorf("result '%s' - expect '%s'", result, expect)
	}
}
