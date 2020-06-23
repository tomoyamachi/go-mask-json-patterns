package sensitive

import (
	"encoding/json"
	"testing"
)

func TestMaskString(t *testing.T) {
	type TestStruct struct {
		Mask   string `sensitive:"true" json:"mask"`
		Normal string `json:"normal" sensitive:"true"`
	}

	tests := []struct {
		in     TestStruct
		expect string
		ok     bool
	}{
		{
			in: TestStruct{
				Mask:   "mask",
				Normal: "normal",
			},
			expect: `{"mask":"***","normal":"nomask"}`,
			ok:     true,
		},
	}
	for i, tt := range tests {
		Mask(tt.in)
		b, err := json.Marshal(tt.in)
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				t.Errorf("test %d, unexpected failure: %v", i, err)
			} else {
				t.Errorf("test %d, unexpected success", i)
			}
		}
		if got := string(b); got != tt.expect {
			t.Errorf("test %d, Marshal(%#v) = %q, want %q", i, tt.in, got, tt.expect)
		}
	}
}
