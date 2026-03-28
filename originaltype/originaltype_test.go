package originaltype

import (
	"encoding/json"
	"testing"

	"github.com/tomoyamachi/go-mask-json-patterns/util"
)

// String()メソッドが常にマスク値を返すことを確認
func TestString(t *testing.T) {
	tests := []struct {
		in     String
		expect string
	}{
		{in: String("hello"), expect: util.Masked},
		{in: String(""), expect: util.Masked},
	}
	for i, tt := range tests {
		got := tt.in.String()
		if got != tt.expect {
			t.Errorf("test %d, String() = %q, want %q", i, got, tt.expect)
		}
	}
}

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
		{
			// 空文字列でもマスクされることを確認
			in: TestStruct{
				Mask:   String(""),
				Normal: "",
			},
			expect: `{"Mask":"***","Normal":""}`,
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
