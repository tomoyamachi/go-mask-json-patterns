package originaltype

import (
	"fmt"
	"github.com/tomoyamachi/go-mask-json-patterns/util"
)

type String string

func (s String) MarshalJSON() ([]byte, error) {
	// Set as json string format
	return []byte(fmt.Sprintf("%q", util.Masked)), nil
}

func (s String) String() string {
	return util.Masked
}
