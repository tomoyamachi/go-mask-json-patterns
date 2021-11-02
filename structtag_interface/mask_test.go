package structtag_interface

import (
	"encoding/json"
	"fmt"
	"github.com/tomoyamachi/go-mask-json-patterns/util"
	"testing"
	"time"
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
