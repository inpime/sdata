package sdata

import (
	"sync"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func TestStringMap_tomldecode(t *testing.T) {
	tomlRaw := `
s = "QWd"


[basic]
tplcache = true

[basic.config]
BrandName = "brand"
BrandCopyright = "copyright"
BrandSite = "http://brand.site"
BrandLogoName = "logo.png" # file name from shop.static, URL("ShopStaticFile", "file", "logo.pnt")
# https://dribbble.com/shots/1518220-Shop-Logo/attachments/229248

[[basic.config.FooterLinks]]
title = "First link"
url = "/first_link"

[[basic.config.FooterLinks]]
title = "Second link"
url = "/second_link"
`
	data := NewStringMap()

	_, err := toml.Decode(tomlRaw, data)
	assert.NoError(t, err)

	if data.M("basic").Bool("tplcache") != true {
		t.FailNow()
	}

	if data.M("basic").M("config").String("BrandName") != "brand" {
		t.FailNow()
	}

	if data.M("basic").M("config").A("FooterLinks").Size() != 2 {
		t.FailNow()
	}

	if NewStringMapFrom(data.M("basic").M("config").A("FooterLinks").Get(0)).String("title") != "First link" {
		t.FailNow()
	}

	if NewStringMapFrom(data.M("basic").M("config").A("FooterLinks").Get(1)).String("title") != "Second link" {
		t.FailNow()
	}
}

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
