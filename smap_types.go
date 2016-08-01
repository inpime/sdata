package sdata

func (m *StringMap) M(key string) *StringMap {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_map := m.getOrNil(key)

	if _map == nil {
		m.set(key, NewStringMap())

		return m.M(key)
	}

	switch _map := _map.(type) {
	case map[string]interface{}:
		m.set(key, NewStringMapFrom(_map))
		return m.M(key)
	case *StringMap:
		return _map
	}

	return nil
}

func (m *StringMap) A(key string) *Array {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_map := m.getOrNil(key)

	if _map == nil {
		m.set(key, NewArray())

		return m.A(key)
	}

	switch _map := _map.(type) {
	case []interface{}:
		m.set(key, NewArray().Add(_map...))
		return m.A(key)
	case *Array:
		return _map
	}

	return nil
}

func (m StringMap) String(key string) string {
	return toString(m.Get(key))
}

func (m StringMap) Float(key string) float64 {
	return toFloat64(m.Get(key))
}

func (m StringMap) Float32(key string) float32 {
	return float32(toFloat64(m.Get(key)))
}

func (m StringMap) Int(key string) int {
	return toInt(m.Get(key))
}

func (m StringMap) Int64(key string) int64 {
	return toInt64(m.Get(key))
}

func (m StringMap) Bool(key string) bool {
	return toBool(m.Get(key))
}

func (m *StringMap) Map(key string) map[string]interface{} {

	if v := m.Get(key); v != nil {
		if v, ok := v.(map[string]interface{}); ok {
			return v
		}
	}

	return map[string]interface{}{}
}
