package main

import (
	"fmt"
	"testing"
)

func TestGreet_English(t *testing.T){
	lang := language("en")
	want := "Hello world!"

	got := greet(lang)

	// fmt.Println(want, got)

	if want != got{
		t.Errorf("expected: %q, got: %q\n", want, got)
	}
}

func TestGreet_French(t *testing.T){
	lang := language("fr")
	want := "Bonjour le monde"

	got := greet(lang)

	// fmt.Println(want, got)
	if want != got{
		t.Errorf("expected: %q, got: %q\n", want, got)
	}
}