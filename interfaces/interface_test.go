package interfaces

import (
	"strings"
	"testing"

	"github.com/tomoyamachi/go-mask-json-patterns/util"
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

// 不正なJSON入力時にエラーを返すことを確認
func TestMaskString_InvalidJSON(t *testing.T) {
	_, err := Log("not json", []string{"foo"})
	if err == nil {
		t.Error("expected error for invalid JSON input, got nil")
	}
}

// 存在しないキーを指定した場合、JSONが変更されないことを確認
func TestMaskString_KeyNotFound(t *testing.T) {
	input := `{"a":"1"}`
	got, err := Log(input, []string{"nonexistent"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ok, err := util.CompareJsonBytes([]byte(got), []byte(input))
	if err != nil {
		t.Fatalf("unexpected error comparing JSON: %v", err)
	}
	if !ok {
		t.Errorf("expected unchanged JSON, got %s", got)
	}
}

// 3階層ネストパスのマスク処理を確認
func TestMaskString_DeepNested(t *testing.T) {
	input := `{"a":{"b":{"c":"secret","d":"keep"}}}`
	expect := `{"a":{"b":{"c":"***","d":"keep"}}}`
	got, err := Log(input, []string{"a/b/c"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ok, err := util.CompareJsonBytes([]byte(got), []byte(expect))
	if err != nil {
		t.Fatalf("unexpected error comparing JSON: %v", err)
	}
	if !ok {
		t.Errorf("got %s, want %s", got, expect)
	}
}

// 中間パスが非オブジェクトの場合、マスクせず値をそのまま置換することを確認
func TestMaskString_NonObjectIntermediate(t *testing.T) {
	input := `{"a":"string_not_object"}`
	got, err := Log(input, []string{"a/b"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 中間のaが文字列なので、aごとマスクされる
	if !strings.Contains(got, `"a":"***"`) {
		t.Errorf("expected a to be masked, got %s", got)
	}
}
