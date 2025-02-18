# go-mask-json-patterns


```go
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
}
type SubMask struct {
Str     string `json:"str,omitempty"`
MaskStr string `json:"mstr,omitempty" log:"*"`
}
```



```go
{
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
  "mpstruct": "*"
}
```
