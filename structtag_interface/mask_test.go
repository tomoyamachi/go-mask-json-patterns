package structtag_interface

import (
	"encoding/json"
	"github.com/tomoyamachi/go-mask-json-patterns/util"
	"testing"
	"time"
)

type MaskResponse struct {
	Str               string            `json:"str,omitempty"`
	MaskStr           string            `json:"mstr,omitempty" log:"*"`
	Int               int               `json:"int,omitempty"`
	MaskInt           int               `json:"mint,omitempty" log:"*"`
	Slice             []string          `json:"slice,omitempty"`
	MaskSlice         []string          `json:"mslice,omitempty" log:"*"`
	Map               map[string]string `json:"map,omitempty"`
	MaskMap           map[string]string `json:"mmap,omitempty" log:"*"`
	Struct            SubMask           `json:"struct,omitempty"`
	MaskStruct        SubMask           `json:"mstruct,omitempty" log:"*"`
	PointerStruct     *SubMask          `json:"pstruct,omitempty"`
	MaskPointerStruct *SubMask          `json:"mpstruct,omitempty" log:"*"`
	Time              time.Time         `json:"time,omitempty"`
	MaskTime          time.Time         `json:"mtime,omitempty" log:"*"`
	PointerTime       *time.Time        `json:"ptime,omitempty"`
	MaskPointerTime   *time.Time        `json:"mptime,omitempty" log:"*"`
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
		Slice:     []string{"a"},
		MaskSlice: []string{"a"},
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
	}
}

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
			in: initMask(),
			ok: true,
			expectLog: `{
  "str": "a",
  "mstr": "*",
  "int": 100,
  "mint": "*",
  "slice": [
    "a"
  ],
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
  "mptime":"*"
}`,
			expectJson: `{
  "str": "a",
  "mstr": "a",
  "int": 100,
  "mint": 100,
  "slice": ["a"],
  "mslice": ["a"],
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
  "mptime":"2000-01-01T00:00:00Z"
}`,
		},
		{
			in: initMaskPtr(),
			ok: true,
			expectLog: `{
  "str": "a",
  "mstr": "*",
  "int": 100,
  "mint": "*",
  "slice": [
    "a"
  ],
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
  "mptime":"*"
}`,
			expectJson: `{
  "str": "a",
  "mstr": "a",
  "int": 100,
  "mint": 100,
  "slice": ["a"],
  "mslice": ["a"],
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
  "mptime":"2000-01-01T00:00:00Z"
}`,
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
