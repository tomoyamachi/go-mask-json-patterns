package override

import (
	"encoding/json"
	"fmt"
	"github.com/tomoyamachi/go-mask-json-patterns/util"
)

type Sample struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type aliasSample Sample

func (s Sample) mask() *aliasSample {
	o := aliasSample(s)
	o.B = util.Masked
	return &o
}

func (s Sample) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.mask())
}

func (s Sample) String() string {
	return fmt.Sprintf("%v", *s.mask())
}
