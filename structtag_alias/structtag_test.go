package structtag_alias

import (
	"encoding/json"
	"github.com/tomoyamachi/go-mask-json-patterns/util"
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
		ok, err := util.CompareJsonBytes(b, []byte(tt.expect))
		if err != nil {
			t.Errorf("test %d, unexpected error with compare log output", i)
		}
		if !ok {
			t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(b), tt.expect)
		}
	}
}
