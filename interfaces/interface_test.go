package interfaces

import (
	"github.com/tomoyamachi/go-mask-json-patterns/util"
	"testing"
)

func TestMaskString(t *testing.T) {
	tests := []struct {
		maskKeys []string
		in       string
		expect   string
		wantErr  error
	}{
		{
			maskKeys: []string{"password"},
			in:       `{"password":"test","user":"tomoyamachi"}`,
			expect:   `{"password":"***","user":"tomoyamachi"}`,
		},
		{
			maskKeys: []string{"password", "authorize_key"},
			in:       `{"authorize_key":"foobar", "password":"test", "user":"tomoyamachi"}`,
			expect:   `{"authorize_key":"***","password":"***","user":"tomoyamachi"}`,
		},
		{
			maskKeys: []string{"password", "authorize_key", "nested/foo"},
			in:       `{"authorize_key":"foobar", "password":"test", "user":"tomoyamachi", "nested": {"foo":123,"bar":456}}`,
			expect:   `{"authorize_key":"***","nested":{"bar":456,"foo":"***"},"password":"***","user":"tomoyamachi"}`,
		},
	}
	for i, tt := range tests {
		got, _ := Log(tt.in, tt.maskKeys)
		ok, err := util.CompareJsonBytes([]byte(got), []byte(tt.expect))
		if err != nil {
			t.Errorf("test %d, unexpected error with compare log output", i)
		}
		if !ok {
			t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(got), tt.expect)
		}
	}
}
