package sdata

func NewArrayFrom(v interface{}) (a *Array) {

	switch v := v.(type) {
	case []interface{}:
		return &Array{data: v}
	case []string:
		a = NewArray()
		for _, value := range v {
			a.Add(value)
		}
	case *Array:
		a = v
	default:
		a = NewArray()
	}

	return
}
