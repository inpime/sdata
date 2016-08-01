package sdata

func (m *Array) Merge(_m Merger) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return Mergex(&m.data, _m.(*Array).data)
}
