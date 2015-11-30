package main

import (
	"os"
	"reflect"
	"testing"
)

func TestIsGo(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	if !IsGo(pwd) {
		t.Errorf("Should be true, but got false")
	}

	if IsGo("/bin") {
		t.Errorf("Should be false, but got true")
	}
}

func TestSplitPATH(t *testing.T) {
	got := SplitPATH("/bin:/usr/bin")

	e := []string{"/bin", "/usr/bin"}
	if !reflect.DeepEqual(e, got) {
		t.Errorf("Expected %q, but got %q", e, got)
	}
}
