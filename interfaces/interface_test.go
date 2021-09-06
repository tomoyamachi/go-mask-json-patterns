package interfaces

import "testing"

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
		if got != tt.expect {
			t.Errorf("test %d, got %q, want %q", i, got, tt.expect)
		}
	}
}
