// file: hello_test.go
package main

import "testing"

func TestSayHello(t *testing.T) {
    got := SayHello("World")
    want := "Hello, World!"
    if got != want {
        t.Errorf("SayHello() = %q, want %q", got, want)
    }
}

