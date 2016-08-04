package sdata

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/BurntSushi/toml"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func NewStringMap() *StringMap {
	return &StringMap{
		data: make(map[string]interface{}),
	}
}

func NewStringMapFrom(v interface{}) (m *StringMap) {
	switch v := v.(type) {
	case map[string]interface{}:
		return &StringMap{
			data: v,
		}
	case *StringMap:
		m = v
	}

	return
}

type StringMap struct {
	mutex sync.RWMutex

	data map[string]interface{}
}

func (m *StringMap) UnmarshalTOML(b []byte) error {
	return toml.Unmarshal(b, m)
}

func (m *StringMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

func (m *StringMap) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.data)
}

func (m *StringMap) MarshalMsgpack() ([]byte, error) {
	return msgpack.Marshal(m.data)
}

func (m *StringMap) UnmarshalMsgpack(b []byte) error {
	dec := msgpack.NewDecoder(bytes.NewBuffer(b))
	dec.DecodeMapFunc = func(d *msgpack.Decoder) (interface{}, error) {
		n, err := d.DecodeMapLen()
		if err != nil {
			return nil, err
		}

		m := make(map[string]interface{}, n)
		for i := 0; i < n; i++ {
			mk, err := d.DecodeString()
			if err != nil {
				return nil, err
			}

			mv, err := d.DecodeInterface()
			if err != nil {
				return nil, err
			}

			m[mk] = mv
		}
		return m, nil
	}

	return dec.Decode(&m.data)
}

func (m *StringMap) Set(key string, value interface{}) *StringMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.set(key, value)

	return m
}

func (m *StringMap) Get(key string) interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.getOrNil(key)
}

func (m *StringMap) Keys() []string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.keys()
}

func (m *StringMap) Remove(key string) *StringMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.remove(key)

	return m
}

// Extend functions

func (m *StringMap) GetIf(key string) (value interface{}, exists bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	value, exists = m.data[key]

	return
}

func (m *StringMap) ToMap() map[string]interface{} {
	out := make(map[string]interface{})

	m.mutex.Lock()
	defer m.mutex.Unlock()

	Mergex(out, m.data)

	return out
}

// private

func (m *StringMap) set(key string, value interface{}) {
	m.data[key] = value
}

func (m *StringMap) get(key string) interface{} {
	return m.data[key]
}

func (m *StringMap) getOrNil(key string) interface{} {
	if value, exists := m.data[key]; exists {
		return value
	}

	return nil
}

func (m *StringMap) size() int {
	return len(m.data)
}

func (m *StringMap) keys() []string {
	keys := make([]string, m.size())
	count := 0
	for key := range m.data {
		keys[count] = key
		count++
	}
	return keys
}

func (m *StringMap) remove(key string) {
	delete(m.data, key)
}
