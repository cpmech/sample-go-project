package count

import "testing"

func TestWords(tst *testing.T) {
	nw := Words("Hello World ")
	if nw != 2 {
		tst.Fail()
	}
}
