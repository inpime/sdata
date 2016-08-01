package sdata

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func TestArr_simple(t *testing.T) {
	a := NewArray()
	a.Add("a")
	a.Add("b", "c").Add("d")
	a.Remove("a")
	a.Add("e")

	assert.Equal(t, a.Size(), 4)
	assert.Equal(t, a.Includes("a"), false)
	assert.Equal(t, a.Includes("e", "b", "c", "d"), true)

	network, err := msgpack.Marshal(a)
	assert.NoError(t, err)

	a = NewArray()

	err = msgpack.Unmarshal(network, a)
	assert.NoError(t, err)

	assert.Equal(t, a.Size(), 4)
	assert.Equal(t, a.Includes("a"), false)
	assert.Equal(t, a.Includes("e", "b", "c", "d"), true)
}

func TestArr_async(t *testing.T) {
	a := NewArray()

	wg := sync.WaitGroup{}

	for k := 0; k < 100; k++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 100; i++ {
				a.Add(i)
			}
			wg.Done()
		}()

		wg.Add(1)

		go func() {
			for i := 0; i < 100; i++ {
				a.Remove(i)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
