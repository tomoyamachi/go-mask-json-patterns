package override

import (
	"encoding/json"
	"fmt"
)

type Sample struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type aliasSample Sample

func (s Sample) mask() *aliasSample {
	o := aliasSample(s)
	o.B = "***"
	return &o
}

func (s Sample) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.mask())
}

func (s Sample) String() string {
	return fmt.Sprintf("%v", *s.mask())
}
