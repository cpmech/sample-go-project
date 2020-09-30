package list

import "testing"

func TestWords(tst *testing.T) {
	words := Words("Hello World ")
	if len(words) != 2 {
		tst.FailNow()
	}
	if words[0] != "Hello" {
		tst.Fail()
	}
	if words[1] != "World" {
		tst.Fail()
	}
}
