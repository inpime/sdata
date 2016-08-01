package sdata

// func TestMerge_sdata_simple(t *testing.T) {
// 	src := NewStringMap()
// 	src.Set("1", "1")
// 	src.Set("2", "2")
// 	src.Set("3", NewStringMap().Set("1", "1").Set("2", "2"))
// 	src.Set("4", NewStringMap().Set("1", "1").Set("2", "2").Set("3", "3"))

// 	dst := NewStringMap()
// 	dst.Set("1", "3")
// 	dst.Set("2", "4")
// 	dst.Set("3", NewStringMap().Set("1", "3").Set("2", "4"))
// 	dst.Set("5", NewStringMap().Set("1", "1").Set("2", "2"))

// 	err := Merge(dst, src)
// 	assert.NoError(t, err)

// 	exp := NewStringMap()
// 	exp.Set("1", "1")
// 	exp.Set("2", "2")
// 	exp.Set("3", NewStringMap().Set("1", "1").Set("2", "2"))
// 	exp.Set("4", NewStringMap().Set("1", "1").Set("2", "2").Set("3", "3"))
// 	exp.Set("5", NewStringMap().Set("1", "1").Set("2", "2"))

// 	t.Logf("%#v", dst.Get("5"))

// 	if !reflect.DeepEqual(dst.data, exp.data) {
// 		t.FailNow()
// 	}
// }

// func TestMerge_sdata_array(t *testing.T) {
// 	src := []interface{}{
// 		NewStringMap().Set("1", "1"),
// 		NewStringMap().Set("2", "2"),
// 		NewStringMap().Set("3", "3"),
// 	}

// 	dst := []interface{}{
// 		NewStringMap().Set("1", "4"),
// 		NewStringMap().Set("2", "5"),
// 		NewStringMap().Set("3", NewStringMap().Set("3", "3")),
// 	}

// 	err := Merge(&dst, src)
// 	assert.NoError(t, err)

// 	exp := []interface{}{
// 		NewStringMap().Set("1", "1"),
// 		NewStringMap().Set("2", "2"),
// 		NewStringMap().Set("3", "3"),
// 		NewStringMap().Set("1", "1"),
// 		NewStringMap().Set("2", "2"),
// 		NewStringMap().Set("3", NewStringMap().Set("3", "3")),
// 	}

// 	if len(dst) != len(exp) {
// 		t.FailNow()
// 	}

// 	// TODO: not a full test
// }

// func TestMerge_sdata_map(t *testing.T) {
// 	src := map[string]interface{}{
// 		"a": NewStringMap().Set("1", "1"),
// 		"b": NewStringMap().Set("2", "2"),
// 		"c": NewStringMap().Set("3", "3"),
// 		"d": NewStringMap().Set("4", NewStringMap().Set("4", "4")),
// 	}

// 	dst := map[string]interface{}{
// 		"a": NewStringMap().Set("1", "11"),
// 		"b": NewStringMap().Set("2", "22"),
// 		"c": NewStringMap().Set("3", NewStringMap().Set("3", "3")),
// 		"e": NewStringMap().Set("5", NewStringMap().Set("5", "5")),
// 	}

// 	err := Merge(dst, src)
// 	assert.NoError(t, err)

// 	exp := map[string]interface{}{
// 		"a": NewStringMap().Set("1", "1"),
// 		"b": NewStringMap().Set("2", "2"),
// 		"c": NewStringMap().Set("3", NewStringMap().Set("3", "3")),
// 		"d": NewStringMap().Set("4", NewStringMap().Set("4", "4")),
// 		"e": NewStringMap().Set("5", NewStringMap().Set("5", "5")),
// 	}

// 	if len(dst) != len(exp) {
// 		t.FailNow()
// 	}

// 	// TODO: not a full test
// }

// type simplestruct struct {
// 	A string
// 	B string
// }

// func TestMerge_simplestruct(t *testing.T) {
// 	src := simplestruct{
// 		A: "a",
// 		B: "b",
// 	}

// 	dst := simplestruct{
// 		A: "a",
// 		B: "c",
// 	}

// 	exp := simplestruct{
// 		A: "a",
// 		B: "b",
// 	}

// 	err := Merge(&dst, src)
// 	assert.NoError(t, err)

// 	if !reflect.DeepEqual(dst, exp) {
// 		t.FailNow()
// 	}
// }

// func TestMerge_simplemap(t *testing.T) {
// 	src := map[string]interface{}{
// 		"a": "a",
// 		"b": "b",
// 	}

// 	dst := map[string]interface{}{
// 		"a": "d",
// 		"b": "e",
// 		"c": "c",
// 	}

// 	exp := map[string]interface{}{
// 		"a": "a",
// 		"b": "b",
// 		"c": "c",
// 	}

// 	err := Merge(&dst, src)
// 	assert.NoError(t, err)

// 	if !reflect.DeepEqual(dst, exp) {
// 		t.FailNow()
// 	}
// }

// type structmap struct {
// 	A map[string]interface{}
// 	B string
// }

// func TestMerge_structmap(t *testing.T) {
// 	src := structmap{
// 		A: map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 			"d": "d",
// 		},
// 		B: "b",
// 	}

// 	dst := structmap{
// 		A: map[string]interface{}{
// 			"a": "d",
// 			"b": "e",
// 			"c": "c",
// 		},
// 		B: "c",
// 	}

// 	exp := structmap{
// 		A: map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 			"c": "c",
// 			"d": "d",
// 		},
// 		B: "b",
// 	}

// 	err := Merge(&dst, src)
// 	assert.NoError(t, err)

// 	if !reflect.DeepEqual(dst, exp) {
// 		t.FailNow()
// 	}
// }

// type structmapstruct struct {
// 	A map[string]simplestruct
// 	B string
// }

// func TestMerge_structmapstruct(t *testing.T) {
// 	src := structmapstruct{
// 		A: map[string]simplestruct{
// 			"a": simplestruct{
// 				A: "aa",
// 				B: "ab",
// 			},
// 			"b": simplestruct{
// 				A: "ba",
// 				B: "bb",
// 			},
// 			"d": simplestruct{
// 				A: "da",
// 				B: "db",
// 			},
// 		},
// 		B: "b",
// 	}

// 	dst := structmapstruct{
// 		A: map[string]simplestruct{
// 			"a": simplestruct{
// 				A: "not valid",
// 				B: "not valid",
// 			},
// 			"b": simplestruct{
// 				A: "not valid",
// 				B: "not valid",
// 			},
// 			"c": simplestruct{
// 				A: "ca",
// 				B: "cb",
// 			},
// 		},
// 		B: "c",
// 	}

// 	exp := structmapstruct{
// 		A: map[string]simplestruct{
// 			"a": simplestruct{
// 				A: "aa",
// 				B: "ab",
// 			},
// 			"b": simplestruct{
// 				A: "ba",
// 				B: "bb",
// 			},
// 			"c": simplestruct{
// 				A: "ca",
// 				B: "cb",
// 			},
// 			"d": simplestruct{
// 				A: "da",
// 				B: "db",
// 			},
// 		},
// 		B: "b",
// 	}

// 	err := Merge(&dst, src)
// 	assert.NoError(t, err)

// 	if !reflect.DeepEqual(dst, exp) {
// 		t.FailNow()
// 	}
// }

// func TestMerge_mapinterface(t *testing.T) {
// 	src := map[string]interface{}{
// 		"a": map[string]interface{}{
// 			"b": "b",
// 		},
// 		"b": map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 		},
// 		"arr":  []string{"a", "b"},
// 		"arr1": []string{"a", "b"},
// 		"arr3": []simplestruct{
// 			simplestruct{"1", "2"},
// 			simplestruct{"3", "4"},
// 		},
// 		"arr4": []map[string]interface{}{
// 			map[string]interface{}{
// 				"1": "2",
// 			},
// 			map[string]interface{}{
// 				"3": "4",
// 			},
// 		},
// 	}

// 	dst := map[string]interface{}{
// 		"a": map[string]interface{}{
// 			"a": "a",
// 			"b": "not valid",
// 		},
// 		"b": map[string]interface{}{
// 			"a": "not valid",
// 			"b": "not valid",
// 		},
// 		"c": map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 		},
// 		"arr":  []string{"c", "d"},
// 		"arr2": []string{"a", "b"},
// 		"arr3": []simplestruct{
// 			simplestruct{"5", "6"},
// 			simplestruct{"7", "8"},
// 		},
// 		"arr4": []map[string]interface{}{
// 			map[string]interface{}{
// 				"5": "6",
// 			},
// 			map[string]interface{}{
// 				"7": "8",
// 			},
// 		},
// 	}

// 	exp := map[string]interface{}{
// 		"a": map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 		},
// 		"b": map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 		},
// 		"c": map[string]interface{}{
// 			"a": "a",
// 			"b": "b",
// 		},
// 		"arr":  []string{"c", "d", "a", "b"},
// 		"arr1": []string{"a", "b"},
// 		"arr2": []string{"a", "b"},
// 		"arr3": []simplestruct{
// 			// important the order of
// 			simplestruct{"5", "6"},
// 			simplestruct{"7", "8"},
// 			simplestruct{"1", "2"},
// 			simplestruct{"3", "4"},
// 		},
// 		"arr4": []map[string]interface{}{
// 			// important the order of
// 			map[string]interface{}{
// 				"5": "6",
// 			},
// 			map[string]interface{}{
// 				"7": "8",
// 			},
// 			map[string]interface{}{
// 				"1": "2",
// 			},
// 			map[string]interface{}{
// 				"3": "4",
// 			},
// 		},
// 	}

// 	err := Merge(&dst, src)
// 	assert.NoError(t, err)

// 	if !reflect.DeepEqual(dst, exp) {
// 		t.FailNow()
// 	}
// }
