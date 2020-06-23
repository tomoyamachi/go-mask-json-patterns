package structtag

import (
	"encoding/json"
	"testing"
)

func TestSensitive(t *testing.T) {

	tests := []struct {
		in     User
		expect string
		ok     bool
	}{
		{
			in: User{
				Id:    1,
				Name:  "tomoya",
				Email: "test.com",
			},
			expect: `{"id":1,"name":"tomoya","email":"***"}`,
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
			t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, got, tt.expect)
		}

	}
}
