# go-mask-json-patterns

A collection of patterns for masking sensitive fields in JSON output with Go.

## Patterns

| Package | Approach | Description |
|---|---|---|
| `structtag_interface` | Struct tag + reflection | Mask fields marked with `log:"*"` tag |
| `structtag_alias` | Struct tag + alias type | Mask fields marked with `sensitive:"true"` tag |
| `interfaces` | JSON key path | Mask fields in a JSON string by specifying key paths |
| `override` | MarshalJSON/String override | Manually implement masking per struct |
| `originaltype` | Custom type | Define a dedicated type that always masks its value |

## Example (`structtag_interface`)

### Struct Definition

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

### Masked Output

```json
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
