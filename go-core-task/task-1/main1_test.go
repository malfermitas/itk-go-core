package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	got := toString(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	want := "42" + "52" + "2a" + "3.14" + "Golang" + "true" + "(1+2i)"
	if got != want {
		t.Errorf("toString() = %q, want %q", got, want)
	}
}

func TestToRunes(t *testing.T) {
	got := toRunes("abc")
	if string(got) != "abc" || len(got) != 3 {
		t.Errorf("toRunes() = %v, want %v", got, []rune{'a', 'b', 'c'})
	}
}

func TestToRunesEmpty(t *testing.T) {
	got := toRunes("")
	if len(got) != 0 {
		t.Errorf("toRunes(\"\") = %v, want empty", got)
	}
}

func TestHashWithSalt(t *testing.T) {
	got := hashWithSalt([]rune("abcdef"))
	if len(got) != 64 {
		t.Errorf("hashWithSalt() len = %d, want 64 (hex sha256)", len(got))
	}

	salted := "abc" + "go-2024" + "def" // соль в середине
	want := fmt.Sprintf("%x", sha256.Sum256([]byte(salted)))
	if got != want {
		t.Errorf("hashWithSalt() = %q, want %q", got, want)
	}
}

func TestHashWithSaltDifference(t *testing.T) {
	a := hashWithSalt([]rune("one"))
	b := hashWithSalt([]rune("two"))
	if a == b {
		t.Error("hashes of different inputs must differ")
	}
}
