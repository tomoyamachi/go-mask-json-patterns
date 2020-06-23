package override

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	tests := []struct {
		in     Sample
		expect string
		ok     bool
	}{
		{
			in: Sample{
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
		fmt.Println(tt.in)
	}
}
