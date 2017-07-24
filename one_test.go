package one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestOneFalse is test for one.gos
func TestOneFalse(t *testing.T) {
	if one(3, 4) != false {
		t.Error("Wrong bool value is coming from one.go")
	}
}

//TestOneTrue will check true being returned
func TestOneTrue(t *testing.T) {
	if one(23, 4) != true {
		t.Error("Wrong bool value is coming from one.go")
	}
}

func TestOneb2(t *testing.T) {
	assert.Equal(t, one(2, 2), true, "this should be true")
}
