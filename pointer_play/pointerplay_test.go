package pointerplay_test

import (
	pointerplay "mymodule/pointer_play"
	"testing"
)

func TestDouble(t *testing.T) {
	t.Parallel()
	x := pointerplay.MyInt(12)
	want := pointerplay.MyInt(24)

	p := &x
	p.Double()
	if want != x {
		t.Errorf("wnat %d, got %d", want, x)
	}
}
