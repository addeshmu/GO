package one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestOneFalse is test for one.go
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

//TestOneb2 tests func two flow 1
func TestOneb2(t *testing.T) {
	assert.Equal(t, one(2, 2), true, "this should be true")
}

//TestTwo1 tests func two flow 2
func TestTwo1(t *testing.T) {
	assert.Equal(t, two(1), false, "false")
	assert.Equal(t, two(10), true, "true")
}
