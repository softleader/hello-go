package slice

import (
	"testing"
)

func TestPlay(t *testing.T) {

	_, m := Play()

	if l := len(m); l != 4 {
		t.Errorf("expected lenth of 4, but got %v", l)
	}
}
