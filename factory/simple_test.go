package factory

import (
	"fmt"
	"testing"
)

var name string = "junbao"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say(name)
	if s != fmt.Sprintf(Hi_FMT, name) {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say(name)
	if s != fmt.Sprintf(HELLO_FMT, name) {
		t.Fatal("Type1 test fail")
	}
}
