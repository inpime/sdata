package sdata

import (
	"sync"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestMergex_structwithstringmap(t *testing.T) {
	type s struct {
		S string
		A *Array
		M *StringMap
	}

	src := &s{
		"a",
		NewArray().Add("a").Add("b"),
		NewStringMap().Set("a", "b").Set("c", "d"),
	}

	dst := &s{
		"b",
		NewArray().Add("c").Add("d"),
		NewStringMap().Set("a", "c").Set("e", "f"),
	}

	exp := &s{
		"a",
		NewArray().Add("a").Add("b").Add("c").Add("d"),
		NewStringMap().Set("a", "b").
			Set("c", "d").
			Set("e", "f"),
	}

	err := Mergex(dst, src)
	assert.NoError(t, err)

	if dst.S != exp.S {
		t.FailNow()
	}

	if !dst.A.Includes(exp.A.Values()...) {
		t.FailNow()
	}

	if !exp.A.Includes(dst.A.Values()...) {
		t.FailNow()
	}

	if !reflect.DeepEqual(exp.M.Data(), dst.M.Data()) {
		t.FailNow()
	}
}

func TestMergex_specialtypes_async(t *testing.T) {
	type s struct {
		S string
		A []string
		I []interface{}
	}

	src := NewStringMap().
		Set("a", "a").
		Set("b", "b").
		Set("c", "c").
		Set("arr", NewArray().Add("a", "b")).
		Set("sc", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		})

	dst := NewStringMap().
		Set("a", "aa").
		Set("b", "bb").
		Set("d", "d").
		Set("arr", NewArray().Add("b", "c")).
		Set("sab", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		}).
		Set("sc", s{
			S: "aa",
			A: []string{"d", "e"},
			I: []interface{}{
				"d",
				[]string{"d", "e"},
			},
		}).
		Set("sd", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		})

	wg := sync.WaitGroup{}

	for k := 0; k < 100; k++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 100; i++ {
				dst.Merge(src)
			}
			wg.Done()
		}()

		wg.Add(1)

		go func() {
			for i := 0; i < 100; i++ {
				dst.Merge(src)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestMergex_specialtypes_merger(t *testing.T) {
	type s struct {
		S string
		A []string
		I []interface{}
	}

	src := NewStringMap().
		Set("a", "a").
		Set("b", "b").
		Set("c", "c").
		Set("arr", NewArray().Add("a", "b")).
		Set("sc", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		})

	dst := NewStringMap().
		Set("a", "aa").
		Set("b", "bb").
		Set("d", "d").
		Set("arr", NewArray().Add("b", "c")).
		Set("sab", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		}).
		Set("sc", s{
			S: "aa",
			A: []string{"d", "e"},
			I: []interface{}{
				"d",
				[]string{"d", "e"},
			},
		}).
		Set("sd", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		})

	exp := NewStringMap().
		Set("a", "a").
		Set("b", "b").
		Set("c", "c").
		Set("d", "d").
		Set("arr", NewArray().Add("a", "b", "b", "c")).
		Set("sab", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		}).
		Set("sc", s{
			S: "a",
			A: []string{"d", "e", "a", "b", "c"},
			I: []interface{}{
				"d",
				[]string{"d", "e"},
				"a",
				[]string{"a", "b", "c"},
			},
		}).
		Set("sd", s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		})

	err := Mergex(dst, src)
	assert.NoError(t, err)

	if dst.Get("a") != exp.Get("a") {
		t.FailNow()
	}

	if dst.Get("b") != exp.Get("b") {
		t.FailNow()
	}

	if dst.Get("c") != exp.Get("c") {
		t.FailNow()
	}

	if dst.Get("d") != exp.Get("d") {
		t.FailNow()
	}

	// arr

	if !dst.A("arr").Includes(exp.A("arr").Values()...) {
		t.FailNow()
	}

	if !exp.A("arr").Includes(dst.A("arr").Values()...) {
		t.FailNow()
	}

	// sab

	if dst.Get("sab").(s).S != exp.Get("sab").(s).S {
		t.FailNow()
	}

	if len(dst.Get("sab").(s).A) != len(exp.Get("sab").(s).A) {
		t.FailNow()
	}

	if len(dst.Get("sab").(s).I) != len(exp.Get("sab").(s).I) {
		t.FailNow()
	}

	// sc

	if dst.Get("sc").(s).S != exp.Get("sc").(s).S {
		t.FailNow()
	}

	if len(dst.Get("sc").(s).A) != len(exp.Get("sc").(s).A) {
		t.FailNow()
	}

	if len(dst.Get("sc").(s).I) != len(exp.Get("sc").(s).I) {
		t.FailNow()
	}

	// sd

	if dst.Get("sd").(s).S != exp.Get("sd").(s).S {
		t.FailNow()
	}

	if len(dst.Get("sd").(s).A) != len(exp.Get("sd").(s).A) {
		t.FailNow()
	}

	if len(dst.Get("sd").(s).I) != len(exp.Get("sd").(s).I) {
		t.FailNow()
	}
}

func TestMergex_struct(t *testing.T) {
	type s struct {
		S string
		A []string
		I []interface{}
	}

	type ss struct {
		S  string
		A  []string
		I  []interface{}
		SS s
	}

	src := ss{
		S: "a",
		A: []string{"a", "b", "c"},
		I: []interface{}{
			"a",
			[]string{"a", "b", "c"},
		},
		SS: s{
			S: "a",
			A: []string{"a", "b", "c"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
			},
		},
	}

	dst := ss{
		S: "aa",
		A: []string{"d", "e"},
		I: []interface{}{
			"d",
			[]string{"e", "f", "g"},
		},
		SS: s{
			S: "aa",
			A: []string{"d", "e"},
			I: []interface{}{
				"d",
				[]string{"e", "f", "g"},
			},
		},
	}

	exp := ss{
		S: "a",
		A: []string{"a", "b", "c", "d", "e"},
		I: []interface{}{
			"a",
			[]string{"a", "b", "c"},
			"d",
			[]string{"e", "f", "g"},
		},
		SS: s{
			S: "a",
			A: []string{"a", "b", "c", "d", "e"},
			I: []interface{}{
				"a",
				[]string{"a", "b", "c"},
				"d",
				[]string{"e", "f", "g"},
			},
		},
	}

	err := Mergex(&dst, src)
	assert.NoError(t, err)

	if dst.S != exp.S {
		t.FailNow()
	}

	if dst.SS.S != exp.SS.S {
		t.FailNow()
	}

	var counter = 0
	for _, dstv := range dst.A {
		for _, expv := range exp.A {
			if dstv == expv {
				counter++
			}
		}
	}

	if len(exp.A) != counter {
		t.FailNow()
	}

	counter = 0
	for _, dstv := range dst.SS.A {
		for _, expv := range exp.SS.A {
			if dstv == expv {
				counter++
			}
		}
	}

	if len(exp.SS.A) != counter {
		t.FailNow()
	}

	if len(exp.I) != len(dst.I) {
		t.FailNow()
	}

	if len(exp.SS.I) != len(dst.SS.I) {
		t.FailNow()
	}
}

func TestMergex_map(t *testing.T) {
	// map

	srcmap := map[string]interface{}{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	dstmap := map[string]interface{}{
		"a": "aa",
		"b": "bb",
		"d": "d",
	}

	expmap := map[string]interface{}{
		"a": "a",
		"b": "b",
		"c": "c",
		"d": "d",
	}

	err := Mergex(&dstmap, srcmap)
	assert.NoError(t, err)

	if !reflect.DeepEqual(dstmap, expmap) {
		t.FailNow()
	}

	// map map

	srcdmap := map[string]interface{}{
		"a": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
		"b": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
		"c": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
	}

	dstdmap := map[string]interface{}{
		"a": map[string]interface{}{
			"a": "aa",
			"b": "bb",
			"c": "cc",
		},
		"b": map[string]interface{}{
			"a": "aa",
			"b": "bb",
			"c": "cc",
		},
		"d": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
	}

	expdmap := map[string]interface{}{
		"a": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
		"b": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
		"c": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
		"d": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
	}

	err = Mergex(&dstdmap, srcdmap)
	assert.NoError(t, err)

	if !reflect.DeepEqual(dstdmap, expdmap) {
		t.FailNow()
	}
}

func TestMergex_slicesftypesandiface(t *testing.T) {

	// arr string

	srcarrstr := []string{"a", "b", "c"}
	dstarrstr := []string{"d", "e"}
	exparrstr := []string{"d", "e", "a", "b", "c"}

	err := Mergex(&dstarrstr, srcarrstr)
	assert.NoError(t, err)

	var counter = 0
	for _, dstv := range dstarrstr {
		for _, expv := range exparrstr {
			if dstv == expv {
				counter++
			}
		}
	}

	if len(exparrstr) != counter {
		t.FailNow()
	}

	// arr interfaces
	srcarriface := []interface{}{"a", "b", "c", 1, 2}
	dstarriface := []interface{}{"d", "e"}
	exparriface := []interface{}{"d", "e", "a", "b", "c", 1, 2}

	err = Mergex(&dstarriface, srcarriface)
	assert.NoError(t, err)

	counter = 0
	for _, dstv := range dstarriface {
		for _, expv := range exparriface {
			if dstv == expv {
				counter++
			}
		}
	}

	if len(exparriface) != counter {
		t.FailNow()
	}
}

func TestMergex_ftypesandiface(t *testing.T) {

	// string

	srcstr := "a"
	dststr := "b"
	expstr := "a"

	err := Mergex(&dststr, srcstr)
	assert.NoError(t, err)

	if dststr != expstr {
		t.FailNow()
	}

	// integer

	srcint := 1
	dstint := 2
	expint := 1

	err = Mergex(&dstint, srcint)
	assert.NoError(t, err)

	if dstint != expint {
		t.FailNow()
	}

	// interface

	var srciface, dstiface, expiface interface{}
	srciface = "a"
	dstiface = "b"
	expiface = "a"

	err = Mergex(&dstiface, srciface)
	assert.NoError(t, err)

	if dstiface != expiface {
		t.FailNow()
	}
}
