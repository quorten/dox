package source_test

import (
	"testing"

	"github.com/jesselang/dox/pkg/source"
)

func TestNew(t *testing.T) {
	_, err := source.New("")
	if err == nil {
		t.Error("empty filename should fail")
	}
}
