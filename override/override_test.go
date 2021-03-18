package override

import (
	"encoding/json"
	"reflect"
	"testing"
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
		if got := string(b); got != tt.expect {
			t.Errorf("test %d, Marshal(%#v) = %q, want %q", i, tt.in, got, tt.expect)
		}
		if !reflect.DeepEqual(tt.in, tt.spare) {
			t.Errorf("test %d, Override original structure", i)
		}
	}
}

func TestSampleToString(t *testing.T) {
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
			expect: `request={"a":"a","b":"***","c":"c"}`,
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
			expect: `request={"a":"a","b":"***","c":"c"}`,
			ok:     true,
		},
	}
	for i, tt := range tests {
		if got := ToString(tt.in); got != tt.expect {
			t.Errorf("test %d, Marshal(%#v) = %q, want %q", i, tt.in, got, tt.expect)
		}
		if !reflect.DeepEqual(tt.in, tt.spare) {
			t.Errorf("test %d, Override original structure", i)
		}
	}
}
