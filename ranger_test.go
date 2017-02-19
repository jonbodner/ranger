package ranger_test

import (
	"testing"
	"github.com/jonbodner/ranger"
)

func TestRange(t *testing.T) {
	vals := make([]bool,10)
	for i := range ranger.UpTo(10) {
		vals[i] = true
	}
	for k, v := range vals {
		if !v {
			t.Errorf("expected %d to be set", k)
		}
	}
}

func TestRangeNegative(t *testing.T) {
	vals := make([]bool,10)
	for i := range ranger.UpTo(-10) {
		vals[i] = true
	}
	for k, v := range vals {
		if v {
			t.Errorf("expected %d to be unset", k)
		}
	}
}

