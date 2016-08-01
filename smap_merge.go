package sdata

import "reflect"

var stringMapType = reflect.TypeOf(&StringMap{})

func (m *StringMap) Merge(_m Merger) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return Mergex(m.data, _m.(*StringMap).data)
}
