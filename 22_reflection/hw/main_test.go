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
		Signature:      "af312b8d12",
	}

	expected := `a=aaaa&g1=12&g2=Tx&g3=SubTx2&s=af312b8d12&v1=42&v2=84`

	result := Serialize(r)
	if result != expected {
		t.Errorf("expected %q, got: %q", expected, result)
	}
}
