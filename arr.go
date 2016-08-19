package sdata

import (
	"encoding/json"
	"sync"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Array struct {
	data  []interface{}
	mutex sync.RWMutex
}

func NewArray() *Array {
	return new(Array)
}

func (a *Array) Add(values ...interface{}) *Array {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.add(values...)

	return a
}

func (a *Array) Get(index int) interface{} {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	if index > a.size() {
		return nil
	}

	return a.get(index)
}

func (a *Array) Remove(value interface{}) *Array {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	i := a.index(value)

	if i == -1 {
		return a
	}

	a.data = a.data[:i+copy((a.data)[i:], (a.data)[i+1:])]

	return a
}

func (a *Array) Index(value interface{}) int {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return a.index(value)
}

func (a *Array) Exist(value interface{}) bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return a.index(value) > -1
}

func (a Array) Includes(values ...interface{}) bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	var count = 0

	for _, value := range values {
		if a.index(value) > -1 {
			count++
		}
	}

	return len(values) == count
}

func (a *Array) Values() []interface{} {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	newArr := make([]interface{}, len(a.data), len(a.data))
	copy(newArr, a.data[:])
	return newArr
}

// Data unsafe function, return data
func (m *Array) Data() []interface{} {
	return m.data
}

func (a *Array) Size() int {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return a.size()
}

func (m Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

func (m *Array) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.data)
}

func (m Array) MarshalMsgpack() ([]byte, error) {
	return msgpack.Marshal(m.data)
}

func (m *Array) UnmarshalMsgpack(b []byte) error {
	return msgpack.Unmarshal(b, &m.data)
}

// private

func (a *Array) get(index int) interface{} {
	return a.data[index]
}

func (a *Array) add(values ...interface{}) {
	a.data = append(a.data, values...)
}

func (a *Array) size() int {

	return len(a.data)
}

func (a *Array) index(value interface{}) int {
	if a.size() == 0 {
		return -1
	}

	for _index, _value := range a.data {
		if _value == value {
			return _index
		}
	}

	return -1
}
