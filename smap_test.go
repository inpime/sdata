package sdata

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func TestSMap_simple(t *testing.T) {
	m := NewStringMap()
	m.Set("a", "b")
	m.Set("c", "d")
	m.M("m").Set("ma", "b")
	m.M("m").Set("mc", "d")

	assert.Equal(t, m.Get("a"), "b")
	assert.Equal(t, m.Get("c"), "d")
	assert.Equal(t, m.M("m").Get("ma"), "b")
	assert.Equal(t, m.M("m").Get("mc"), "d")

	network, err := msgpack.Marshal(m)
	assert.NoError(t, err)

	m = NewStringMap()

	err = msgpack.Unmarshal(network, m)
	assert.NoError(t, err)

	assert.Equal(t, m.Get("a"), "b")
	assert.Equal(t, m.Get("c"), "d")
	assert.Equal(t, m.M("m").Get("ma"), "b")
	assert.Equal(t, m.M("m").Get("mc"), "d")
}

func TestSMap_async(t *testing.T) {
	m := NewStringMap()
	m.Set("a", "b")
	m.Set("c", "d")
	m.M("m").Set("v", 0)

	wg := sync.WaitGroup{}

	for k := 0; k < 100; k++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				m.M("m").Set("v", i)
			}
			wg.Done()
		}()

		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				m.M("m").Get("v")
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
