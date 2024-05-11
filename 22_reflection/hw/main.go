package main

import "fmt"

type Request struct {
	Alpha          string `param:"a"`
	GroupIndex     int    `param:"g1"`
	GroupType      string `param:"g2"`
	GroupSubtype   string `param:"g3"`
	ValuePrimary   int64  `param:"v1"`
	ValueSecondary int64  `param:"v2"`

	Signature string `param:"s"`
}

func main() {
	r := Request{
		Alpha:          "aaaa",
		GroupIndex:     12,
		GroupType:      "Tx",
		GroupSubtype:   "SubTx2",
		ValuePrimary:   42,
		ValueSecondary: 84,
		Signature:      "af312b8d12",
	}

	fmt.Println(Serialize(r))
}
