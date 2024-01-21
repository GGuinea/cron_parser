package parser

import (
	"testing"
)

func TestShouldReturnNewEmptyParserClient(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Fatalf("Should return new parser")
	}
}
