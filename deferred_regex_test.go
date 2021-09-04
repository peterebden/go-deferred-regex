package deferredregex

import (
	"testing"
)

func TestReplacement(t *testing.T) {
	dr := DeferredRegex{Re: `([0-9]+)\.([0-9]+)\.([0-9]+)`}
	matches := dr.FindStringSubmatch("1.2.3")
	if len(matches) != 4 || matches[1] != "1" || matches[2] != "2" || matches[3] != "3" {
		t.Errorf("Failed to match string")
	}
}
