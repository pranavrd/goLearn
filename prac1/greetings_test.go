package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Raymond"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := HelloOld("Raymond")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Raymond") = %q, %v, want match for %#q, nil`, msg, err, want)

	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := HelloOld("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")=%q,%v,want "", error`, msg, err)
	}
}
