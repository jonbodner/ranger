package ranger_test

import (
	"testing"
	"github.com/jonbodner/ranger"
	"fmt"
)

func TestRange(t *testing.T) {
	for i := range ranger.UpTo(10) {
		fmt.Println(i)
	}
}


