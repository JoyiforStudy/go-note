package main

import "testing"

const englishHelloPrefix = "hello, "

func hello(str string) string {
	return englishHelloPrefix + str
}

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("fciasth")
		get := "hello, fciasth"

		if got != get {
			t.Errorf("got '%q' get '%q'", got, get)
		}
	})
}
