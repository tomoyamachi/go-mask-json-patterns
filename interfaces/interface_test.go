package interfaces

import "testing"

func TestMaskString(t *testing.T) {
	tests := []struct {
		in      string
		expect  string
		wantErr error
	}{
		{
			in:     `{"password":"test","user":"tomoyamachi"}`,
			expect: `{"password":"***","user":"tomoyamachi"}`,
		},
		{
			in:     `{"authorize_key":"foobar", "password":"test", "user":"tomoyamachi"}`,
			expect: `{"authorize_key":"***","password":"***","user":"tomoyamachi"}`,
		},
	}
	for i, tt := range tests {
		got, _ := Log(tt.in)
		if got != tt.expect {
			t.Errorf("test %d, input(%#v) = %q, want %q", i, tt.in, got, tt.expect)
		}
	}
}
