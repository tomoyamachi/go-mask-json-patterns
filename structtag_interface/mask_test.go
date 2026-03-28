package structtag_interface

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/tomoyamachi/go-mask-json-patterns/util"
)

type MaskResponse struct {
	Str                 string            `json:"str,omitempty"`
	MaskStr             string            `json:"mstr,omitempty" log:"*"`
	Int                 int               `json:"int,omitempty"`
	MaskInt             int               `json:"mint,omitempty" log:"*"`
	Slice               []string          `json:"slice,omitempty"`
	MaskSlice           []string          `json:"mslice,omitempty" log:"*"`
	Map                 map[string]string `json:"map,omitempty"`
	MaskMap             map[string]string `json:"mmap,omitempty" log:"*"`
	Struct              SubMask           `json:"struct,omitempty"`
	MaskStruct          SubMask           `json:"mstruct,omitempty" log:"*"`
	PointerStruct       *SubMask          `json:"pstruct,omitempty"`
	MaskPointerStruct   *SubMask          `json:"mpstruct,omitempty" log:"*"`
	Time                time.Time         `json:"time,omitempty"`
	MaskTime            time.Time         `json:"mtime,omitempty" log:"*"`
	PointerTime         *time.Time        `json:"ptime,omitempty"`
	MaskPointerTime     *time.Time        `json:"mptime,omitempty" log:"*"`
	SubMasks            []SubMask         `json:"structs,omitempty"`
	MaskSubMasks        []SubMask         `json:"mstructs,omitempty" log:"*"`
	PointerSubMasks     []*SubMask        `json:"pstructs,omitempty"`
	MaskPointerSubMasks []*SubMask        `json:"mpstructs,omitempty" log:"*"`
}

type SubMask struct {
	Str     string `json:"str,omitempty"`
	MaskStr string `json:"mstr,omitempty" log:"*"`
}

var dummyTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

func initMask() MaskResponse {
	return MaskResponse{
		Str:       "a",
		MaskStr:   "a",
		Int:       100,
		MaskInt:   100,
		Slice:     []string{"a", "b"},
		MaskSlice: []string{"a", "b"},
		Map:       map[string]string{"a": "b"},
		MaskMap:   map[string]string{"a": "b"},
		Struct: SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		MaskStruct: SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		PointerStruct: &SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		MaskPointerStruct: &SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		Time:            dummyTime,
		MaskTime:        dummyTime,
		PointerTime:     &dummyTime,
		MaskPointerTime: &dummyTime,
		SubMasks: []SubMask{
			{Str: "a", MaskStr: "a"},
			{Str: "b", MaskStr: "b"},
		},
		MaskSubMasks: []SubMask{
			{Str: "a", MaskStr: "a"},
			{Str: "b", MaskStr: "b"},
		},
		PointerSubMasks: []*SubMask{
			{Str: "a", MaskStr: "a"},
			{Str: "b", MaskStr: "b"},
		},
		MaskPointerSubMasks: []*SubMask{
			{Str: "a", MaskStr: "a"},
			{Str: "b", MaskStr: "b"},
		},
	}
}

var expectMaskedStringInitMask = `{
  "str": "a",
  "mstr": "*",
  "int": 100,
  "mint": "*",
  "slice": ["a", "b"],
  "mslice": "*",
  "map": {
    "a": "b"
  },
  "mmap": "*",
  "struct": {
    "mstr": "*",
    "str": "a"
  },
  "mstruct": "*",
  "pstruct": {
    "mstr": "*",
    "str": "a"
  },
  "mpstruct": "*",
  "time":{"v":"2000-01-01T00:00:00Z"},
  "mtime":"*",
  "ptime":{"v":"2000-01-01T00:00:00Z"},
  "mptime":"*",
  "structs": [{"mstr": "*","str": "a"},{"mstr": "*","str": "b"}],
  "mstructs": "*",
  "pstructs": [{"mstr": "*","str": "a"},{"mstr": "*","str": "b"}],
  "mpstructs": "*"
}`

var expectStringInitMask = `{
  "str": "a",
  "mstr": "a",
  "int": 100,
  "mint": 100,
  "slice": ["a","b"],
  "mslice": ["a","b"],
  "map": {"a": "b"},
  "mmap": {"a": "b"},
  "struct": {
    "mstr": "a",
    "str": "a"
  },
  "mstruct": {
    "mstr": "a",
    "str": "a"
  },
  "pstruct": {
    "mstr": "a",
    "str": "a"
  },
  "mpstruct": {
    "mstr": "a",
    "str": "a"
  },
  "time":"2000-01-01T00:00:00Z",
  "mtime":"2000-01-01T00:00:00Z",
  "ptime":"2000-01-01T00:00:00Z",
  "mptime":"2000-01-01T00:00:00Z",
  "structs": [{"mstr": "a","str": "a"},{"mstr": "b","str": "b"}],
  "mstructs": [{"mstr": "a","str": "a"},{"mstr": "b","str": "b"}],
  "pstructs": [{"mstr": "a","str": "a"},{"mstr": "b","str": "b"}],
  "mpstructs": [{"mstr": "a","str": "a"},{"mstr": "b","str": "b"}]
}`

func initMaskPtr() *MaskResponse {
	m := initMask()
	return &m
}

func TestMask(t *testing.T) {
	pointerStr := "pointer str"
	tests := []struct {
		in         interface{}
		ok         bool
		expectLog  string
		expectJson string
	}{
		{
			in:         initMask(),
			ok:         true,
			expectLog:  expectMaskedStringInitMask,
			expectJson: expectStringInitMask,
		},
		{
			in:         initMaskPtr(),
			ok:         true,
			expectLog:  expectMaskedStringInitMask,
			expectJson: expectStringInitMask,
		},
		{
			in:         []MaskResponse{initMask(), initMask()},
			ok:         true,
			expectLog:  fmt.Sprintf(`{"masked slice":[%s,%s]}`, expectMaskedStringInitMask, expectMaskedStringInitMask),
			expectJson: fmt.Sprintf("[%s,%s]", expectStringInitMask, expectStringInitMask),
		},
		{
			in:         "normal str",
			ok:         true,
			expectLog:  `{"msg":"normal str"}`,
			expectJson: `"normal str"`,
		},
		{
			in:         &pointerStr,
			ok:         true,
			expectLog:  `{"msg":"pointer str"}`,
			expectJson: `"pointer str"`,
		},
		{
			in:         nil,
			ok:         true,
			expectLog:  `null`,
			expectJson: `null`,
		},
	}
	for i, tt := range tests {
		b, err := Log(tt.in)
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				t.Errorf("test %d, unexpected failure: %v", i, err)
			} else {
				t.Errorf("test %d, unexpected success", i)
			}
			continue
		}

		// check log output
		{
			ok, err := util.CompareJsonBytes(b, []byte(tt.expectLog))
			if err != nil {
				t.Errorf("test %d, unexpected error with compare log output", i)
			}
			if !ok {
				t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(b), tt.expectLog)
			}
		}

		// check json marshal
		got, err := json.Marshal(tt.in)
		if err != nil {
			t.Errorf("test %d, failed json.Marshal", i)
		}
		{
			ok, err := util.CompareJsonBytes(got, []byte(tt.expectJson))
			if err != nil {
				t.Errorf("test %d, unexpected error with compare log output", i)
			}
			if !ok {
				t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(got), tt.expectJson)
			}
		}

	}
}

// 非公開フィールド判定のテスト
func TestIsPrivateField(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
	}{
		{name: "Name", expect: false},
		{name: "name", expect: true},
		{name: "", expect: true},
		{name: "X", expect: false},
		{name: "x", expect: true},
	}
	for i, tt := range tests {
		got := isPrivateField(tt.name)
		if got != tt.expect {
			t.Errorf("test %d, isPrivateField(%q) = %v, want %v", i, tt.name, got, tt.expect)
		}
	}
}

// jsonタグ解析のテスト
func TestParseTag(t *testing.T) {
	tests := []struct {
		tag        string
		expectName string
		expectOpts tagOptions
	}{
		{tag: "name,omitempty", expectName: "name", expectOpts: "omitempty"},
		{tag: "name", expectName: "name", expectOpts: ""},
		{tag: ",omitempty", expectName: "", expectOpts: "omitempty"},
		{tag: "", expectName: "", expectOpts: ""},
	}
	for i, tt := range tests {
		name, opts := parseTag(tt.tag)
		if name != tt.expectName {
			t.Errorf("test %d, parseTag(%q) name = %q, want %q", i, tt.tag, name, tt.expectName)
		}
		if opts != tt.expectOpts {
			t.Errorf("test %d, parseTag(%q) opts = %q, want %q", i, tt.tag, opts, tt.expectOpts)
		}
	}
}

// tagOptions.Containsのテスト
func TestTagOptions_Contains(t *testing.T) {
	tests := []struct {
		opts   tagOptions
		name   string
		expect bool
	}{
		{opts: "omitempty,string", name: "omitempty", expect: true},
		{opts: "omitempty,string", name: "string", expect: true},
		{opts: "omitempty,string", name: "missing", expect: false},
		{opts: "", name: "anything", expect: false},
		{opts: "omitempty", name: "omitempty", expect: true},
	}
	for i, tt := range tests {
		got := tt.opts.Contains(tt.name)
		if got != tt.expect {
			t.Errorf("test %d, tagOptions(%q).Contains(%q) = %v, want %v", i, tt.opts, tt.name, got, tt.expect)
		}
	}
}

// 各型のゼロ値判定テスト
func TestIsEmptyValue(t *testing.T) {
	var nilPtr *int
	nonNilPtr := new(int)
	tests := []struct {
		desc   string
		val    any
		expect bool
	}{
		{desc: "empty string", val: "", expect: true},
		{desc: "non-empty string", val: "hello", expect: false},
		{desc: "zero int", val: 0, expect: true},
		{desc: "non-zero int", val: 42, expect: false},
		{desc: "false bool", val: false, expect: true},
		{desc: "true bool", val: true, expect: false},
		{desc: "zero float", val: 0.0, expect: true},
		{desc: "non-zero float", val: 1.5, expect: false},
		{desc: "nil pointer", val: nilPtr, expect: true},
		{desc: "non-nil pointer", val: nonNilPtr, expect: false},
		{desc: "empty slice", val: []string{}, expect: true},
		{desc: "non-empty slice", val: []string{"a"}, expect: false},
	}
	for _, tt := range tests {
		got := isEmptyValue(reflect.ValueOf(tt.val))
		if got != tt.expect {
			t.Errorf("isEmptyValue(%s) = %v, want %v", tt.desc, got, tt.expect)
		}
	}
}

// checkSpecialStructのテスト
func TestCheckSpecialStruct(t *testing.T) {
	// time.Timeは特殊構造体として元の値を返す
	tv := time.Now()
	rv := reflect.ValueOf(tv)
	got := checkSpecialStruct(rv)
	if got == nil {
		t.Error("checkSpecialStruct(time.Time) should return non-nil")
	}

	// 通常の構造体はnilを返す
	sv := SubMask{Str: "a"}
	rv2 := reflect.ValueOf(sv)
	got2 := checkSpecialStruct(rv2)
	if got2 != nil {
		t.Errorf("checkSpecialStruct(SubMask) should return nil, got %v", got2)
	}
}

// jsonタグなし構造体のテスト（フィールド名がキーとして使用される）
func TestMakeMaskedStruct_NoJsonTags(t *testing.T) {
	type NoTag struct {
		Name   string
		Secret string `log:"*"`
	}
	input := NoTag{Name: "visible", Secret: "hidden"}
	result := MakeMaskedStruct(input)

	b, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `{"Name":"visible","Secret":"*"}`
	ok, err := util.CompareJsonBytes(b, []byte(expect))
	if err != nil {
		t.Fatalf("unexpected error comparing JSON: %v", err)
	}
	if !ok {
		t.Errorf("got %s, want %s", string(b), expect)
	}
}
