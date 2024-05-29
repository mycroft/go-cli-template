package internal

import "testing"

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "hello world!" {
		t.Fatal("invalid value returned")
	}
}
