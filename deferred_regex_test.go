package deferredregex

import (
	"encoding"
	"testing"
)

func TestReplacement(t *testing.T) {
	dr := DeferredRegex{Re: `([0-9]+)\.([0-9]+)\.([0-9]+)`}
	matches := dr.FindStringSubmatch("1.2.3")
	if len(matches) != 4 || matches[1] != "1" || matches[2] != "2" || matches[3] != "3" {
		t.Errorf("Failed to match string")
	}
}

func TestTextUnmarshaler(t *testing.T) {
	const ex = `([0-9]+)\.([0-9]+)\.([0-9]+)`
	var m encoding.TextMarshaler = &DeferredRegex{Re: ex}
	text, err := m.MarshalText()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	} else if string(text) != ex {
		t.Errorf("mismatching regexes")
	}
	var um encoding.TextUnmarshaler = &DeferredRegex{}
	err = um.UnmarshalText(text)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestFlagsMarshaler(t *testing.T) {
	type FlagMarshaler interface {
		UnmarshalFlag(value string) error
		MarshalFlag() (string, error)
	}

	const ex = `([0-9]+)\.([0-9]+)\.([0-9]+)`
	var m FlagMarshaler = &DeferredRegex{Re: ex}
	text, err := m.MarshalFlag()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	} else if text != ex {
		t.Errorf("mismatching regexes")
	}
	err = m.UnmarshalFlag(text)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}
