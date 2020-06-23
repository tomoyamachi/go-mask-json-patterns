package originaltype

type String string

func (s String) MarshalJSON() ([]byte, error) {
	return []byte(`"***"`), nil
}

func (s String) String() string {
	return "***"
}
