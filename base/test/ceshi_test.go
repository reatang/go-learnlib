package test

import "testing"

func TestSayHello(t *testing.T) {
	talk := SayHello()

	if talk != "hello" {
		t.Errorf("SayHello not say 'hello'")
	}
}
