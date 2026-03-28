package override

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/tomoyamachi/go-mask-json-patterns/util"
)

func TestSample(t *testing.T) {
	tests := []struct {
		in     interface{}
		spare  interface{}
		expect string
		ok     bool
	}{
		{
			in: Sample{
				A: "a",
				B: "b",
				C: "c",
			},
			spare: Sample{
				A: "a",
				B: "b",
				C: "c",
			},
			expect: `{"a":"a","b":"***","c":"c"}`,
			ok:     true,
		},
		{
			in: &Sample{
				A: "a",
				B: "b",
				C: "c",
			},
			spare: &Sample{
				A: "a",
				B: "b",
				C: "c",
			},
			expect: `{"a":"a","b":"***","c":"c"}`,
			ok:     true,
		},
		{
			// 空構造体でもBがマスクされることを確認
			in:     Sample{},
			spare:  Sample{},
			expect: `{"a":"","b":"***","c":""}`,
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
		// check original value does not change
		if !reflect.DeepEqual(tt.in, tt.spare) {
			t.Errorf("test %d, Override original structure", i)
		}
	}
}

// String()メソッドがBフィールドをマスクすることを確認
func TestSample_String(t *testing.T) {
	s := Sample{A: "a", B: "b", C: "c"}
	got := s.String()
	if strings.Contains(got, "b") && !strings.Contains(got, util.Masked) {
		t.Errorf("String() should mask field B, got %s", got)
	}
	if !strings.Contains(got, util.Masked) {
		t.Errorf("String() should contain masked value, got %s", got)
	}
	// 元の構造体が変更されていないことを確認
	if s.B != "b" {
		t.Errorf("original struct was mutated: B = %q", s.B)
	}
}
