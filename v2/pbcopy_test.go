package main

import (
	"runtime"
	"testing"
)

func TestPBcopy(t *testing.T) {
	if runtime.GOOS != "darwin" {
		return
	}
	const oString = "ost"
	originalString, err := pbpaste()
	if err != nil {
		t.Fail()
	}
	if err := pbcopy(oString); err != nil {
		t.Fail()
	}
	if s, err := pbpaste(); err != nil {
		t.Fail()
	} else {
		if s != oString {
			t.Fail()
		}
	}
	if err = pbcopy(originalString); err != nil {
		t.Fail()
	}
}
