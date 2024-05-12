package main

import "testing"

func TestSerialize(t *testing.T) {
	r := Request{
		Alpha:          "aaaa",
		GroupIndex:     12,
		GroupType:      "Tx",
		GroupSubtype:   "SubTx2",
		ValuePrimary:   42,
		ValueSecondary: 84,
		Timestamp:      1234567890,
	}

	expResult := `a=aaaa g1=12 g2=Tx t=1234567890 v1=42 v2=84`
	result := Serialize(r)
	if result != expResult {
		t.Errorf("expected %q, got: %q", expResult, result)
	}
}
