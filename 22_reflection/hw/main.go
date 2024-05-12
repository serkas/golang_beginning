package main

import (
	"fmt"
	"time"
)

type Request struct {
	Alpha          string `param:"a"`
	GroupIndex     int    `param:"g1"`
	GroupType      string `param:"g2"`
	GroupSubtype   string `param:"-"`
	ValuePrimary   int64  `param:"v1"`
	ValueSecondary int64  `param:"v2"`

	Timestamp int64 `param:"t"`
}

func main() {
	r := Request{
		Alpha:          "aaaa",
		GroupIndex:     12,
		GroupType:      "Tx",
		GroupSubtype:   "SubTx2",
		ValuePrimary:   42,
		ValueSecondary: 84,
		Timestamp:      time.Now().Unix(),
	}

	fmt.Println(Serialize(r))
}
