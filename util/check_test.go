package util

import "testing"

func TestCompareJsonBytes(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		hasErr bool
		ok     bool
	}{
		{
			a:  `{"foo":"1","bar":"1"}`,
			b:  `{"bar":"1","foo":"1"}`,
			ok: true,
		},
		{
			// https://json.org/example.html
			a: `{"glossary": {
    "GlossDiv": {
      "title": "S",
      "GlossList": {
        "GlossEntry": {
          "ID": "SGML",
          "SortAs": "SGML",
          "GlossTerm": "Standard Generalized Markup Language",
          "Acronym": "SGML",
          "Abbrev": "ISO 8879:1986",
          "GlossDef": {
            "para": "A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso": [
              "GML","XML"
            ]
          },
          "GlossSee": "markup"
        }
      }
    },
    "title": "example glossary"
  }
}`,
			b:  `{"glossary":{"title":"example glossary","GlossDiv":{"title":"S","GlossList":{"GlossEntry":{"ID":"SGML","SortAs":"SGML","GlossTerm":"Standard Generalized Markup Language","Acronym":"SGML","Abbrev":"ISO 8879:1986","GlossDef":{"para":"A meta-markup language, used to create markup languages such as DocBook.","GlossSeeAlso":["GML","XML"]},"GlossSee":"markup"}}}}}`,
			ok: true,
		},
		{
			// detect slice order change
			a:  `{"foo":[1,2,3]}`,
			b:  `{"foo":[3,2,1]}`,
			ok: false,
		},
		{
			a:      `invalid value`,
			b:      `{"bar":"1","foo":"1"}`,
			hasErr: true,
		},
	}
	for i, tt := range tests {
		ok, err := CompareJsonBytes([]byte(tt.a), []byte(tt.b))
		if err != nil && !tt.hasErr {
			t.Errorf("test %d, unexpected error occured %v", i, err)
			continue
		}
		if ok != tt.ok {
			t.Errorf("test %d, %s and %s", i, tt.a, tt.b)
		}
	}
}
