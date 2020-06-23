package mask

type String string

func (s String) MarshalJSON() ([]byte, error) {
	return []byte(`"***"`), nil
}
