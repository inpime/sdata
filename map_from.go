package sdata

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
)

func (m *Map) LoadFrom(v interface{}) *Map {
	switch v := v.(type) {
	case string:
		m.mutex.Lock()
		defer m.mutex.Unlock()

		if len(v) > 0 {
			err := json.Unmarshal([]byte(v), m)
			if err != nil {
				logrus.WithField("_service", "utils").WithError(err).Error("map: load from json")
			}
		}
	case map[string]interface{}:
		m.LoadFromMapStrIface(v)
	case map[interface{}]interface{}:
		m.LoadFromMapIfaceIface(v)
	case *Map:
		m.LoadFromMap(v)
	}

	return m
}

func (m *Map) LoadFromMapStrIface(v map[string]interface{}) *Map {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for key, value := range v {
		m.set(key, value)
	}
	return m
}

func (m *Map) LoadFromMap(v *Map) *Map {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, key := range v.keys() {
		m.set(key, v.getOrNil(key))
	}
	return m
}

func (m *Map) LoadFromMapIfaceIface(v map[interface{}]interface{}) *Map {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for key, value := range v {
		m.set(key, value)
	}
	return m
}

func (m *Map) Clear() *Map {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, key := range m.Keys() {
		m.remove(key)
	}

	return m
}
