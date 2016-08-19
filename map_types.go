package sdata

func (m *Map) M(key interface{}) *Map {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_map := m.getOrNil(key)

	if _map == nil {
		m.set(key, NewMap())

		return m.M(key)
	}

	switch _map := _map.(type) {
	case map[interface{}]interface{}:
		m.set(key, NewMapFrom(_map))
		return m.M(key)
	case *Map:
		return _map
	}

	return nil
}

func (m *Map) A(key interface{}) *Array {
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

func (m Map) String(key interface{}) string {
	return toString(m.Get(key))
}

func (m Map) Float(key interface{}) float64 {
	return toFloat64(m.Get(key))
}

func (m Map) Float32(key interface{}) float32 {
	return float32(toFloat64(m.Get(key)))
}

func (m Map) Int(key interface{}) int {
	return toInt(m.Get(key))
}

func (m Map) Int64(key interface{}) int64 {
	return toInt64(m.Get(key))
}

func (m Map) Bool(key interface{}) bool {
	return toBool(m.Get(key))
}

func (m *Map) Map(key interface{}) map[interface{}]interface{} {

	if v := m.Get(key); v != nil {
		if v, ok := v.(map[interface{}]interface{}); ok {
			return v
		}
	}

	return map[interface{}]interface{}{}
}
