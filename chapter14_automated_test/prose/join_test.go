package prose

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("did't match expected value")
	}
}
