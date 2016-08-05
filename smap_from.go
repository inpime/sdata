package sdata

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
)

func (m *StringMap) LoadFrom(v interface{}) *StringMap {
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
	case *StringMap:
		m.LoadFromStringMap(v)
	}

	return m
}

func (m *StringMap) LoadFromMapStrIface(v map[string]interface{}) *StringMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for key, value := range v {
		m.set(key, value)
	}
	return m
}

func (m *StringMap) LoadFromStringMap(v *StringMap) *StringMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, key := range v.keys() {
		m.set(key, v.getOrNil(key))
	}
	return m
}

func (m *StringMap) LoadFromMapIfaceIface(v map[interface{}]interface{}) *StringMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for key, value := range v {
		m.set(toString(key), value)
	}
	return m
}

func (m *StringMap) Clear() *StringMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, key := range m.keys() {
		m.remove(key)
	}

	return m
}
