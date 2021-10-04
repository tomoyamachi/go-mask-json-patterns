package originaltype

import (
	"encoding/json"
	"github.com/tomoyamachi/go-mask-json-patterns/util"
	"testing"
)

func TestMaskString(t *testing.T) {
	type TestStruct struct {
		Mask   String `json:"Mask"`
		Normal string `json:"Normal"`
	}

	tests := []struct {
		in     TestStruct
		expect string
		ok     bool
	}{
		{
			in: TestStruct{
				Mask:   String("Mask"),
				Normal: "nomask",
			},
			expect: `{"Mask":"***","Normal":"nomask"}`,
			ok:     true,
		},
	}
	for i, tt := range tests {
		b, err := json.Marshal(tt.in)
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				t.Errorf("test %d, unexpected failure: %v", i, err)
			} else {
				t.Errorf("test %d, unexpected success", i)
			}
		}

		ok, err := util.CompareJsonBytes(b, []byte(tt.expect))
		if err != nil {
			t.Errorf("test %d, unexpected error with compare log output", i)
		}
		if !ok {
			t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(b), tt.expect)
		}
	}
}
