package sdata

import (
	"fmt"
	"reflect"
)

// Merger
type Merger interface {
	Merge(Merger) error
}

func Merge(dst, src interface{}) error {
	return mergex(reflect.ValueOf(dst), reflect.ValueOf(src))
}

func Mergex(dst, src interface{}) error {
	return mergex(reflect.ValueOf(dst), reflect.ValueOf(src))
}

func mergex(vdst, vsrc reflect.Value) error {

	//
	// special type
	//

	if _, special, err := mergeIfSpecialType(vdst, vsrc); special {

		return err
	}

	if !vsrc.CanInterface() {
		return fmt.Errorf("vsrc can not interface")
	}

	if vdst.Kind() == reflect.Interface {
		if !vdst.CanInterface() {
			return fmt.Errorf("vdst can not interface")
		}

		return mergex(reflect.ValueOf(vdst.Interface()), vsrc)
	}

	//
	// simple type
	//

	vdst = reflect.Indirect(vdst)
	vsrc = reflect.Indirect(vsrc)

	switch vdst.Kind() {
	case reflect.Struct:
		for i := 0; i < vdst.NumField(); i++ {
			if err := mergex(vdst.Field(i), vsrc.Field(i)); err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i := 0; i < vsrc.Len(); i++ {
			vdst.Set(reflect.Append(vdst, vsrc.Index(i)))
		}
	case reflect.Map:
		for _, key := range vsrc.MapKeys() {
			vsrcElem := vsrc.MapIndex(key)
			if err := dstAsMap(key, vdst, vsrcElem); err != nil {
				return err
			}
		}
	default:
		if vdst.CanSet() {
			vdst.Set(vsrc)
		}
	}

	return nil
}

// mergeIfSpecialType
func mergeIfSpecialType(vdstIn, vsrc reflect.Value) (vdst reflect.Value, isSpecial bool, err error) {
	vdst = vdstIn

	switch vsrc.Interface().(type) {
	case Merger:
		isSpecial = true

		switch vsrc.Interface().(type) {
		case *StringMap:
			if !vdstIn.IsValid() {
				// TODO: universal factory special type object
				vdst = reflect.ValueOf(NewStringMap())
			}
		case *Array:
			if !vdstIn.IsValid() {
				// TODO: universal factory special type object
				vdst = reflect.ValueOf(NewArray())
			}
		default:
			err = fmt.Errorf("not supported type %T", vsrc.Interface())
			return
		}

		// TODO: check if diff types

		err = vdst.Interface().(Merger).Merge(vsrc.Interface().(Merger))
	default:
		isSpecial = false
	}

	return
}

func dstAsMap(key, vdst, vsrc reflect.Value) error {
	// vsrc = reflect.ValueOf(vsrc.Interface())

	if !vsrc.CanInterface() {
		return fmt.Errorf("vdst element %q can not interface", key.String())
	}

	//
	// special type
	//

	if _vdst, special, err := mergeIfSpecialType(vdst.MapIndex(key), vsrc); special {
		vdst.SetMapIndex(key, _vdst)

		return err
	}

	//
	// simple type
	//

	switch vsrc.Kind() {
	case reflect.Interface:
		return dstAsMap(key, vdst, reflect.ValueOf(vsrc.Interface()))
	case reflect.Map, reflect.Struct, reflect.Slice:
		vdstElem := vdst.MapIndex(key)

		if !vdstElem.CanSet() {
			vdstElem = reflect.New(vsrc.Type()).Elem()

			if vsrc.Kind() == reflect.Map {
				vdstElem.Set(reflect.MakeMap(vsrc.Type()))
			}

			if vdstElem.IsValid() {

				if vdst.MapIndex(key).IsValid() {
					// default from src
					vdstElem.Set(reflect.ValueOf(vdst.MapIndex(key).Interface()))
				}
			}
		}

		if err := mergex(vdstElem, reflect.Indirect(vsrc)); err != nil {
			return err
		}

		vdst.SetMapIndex(key, reflect.Indirect(vdstElem))
	default:
		vdst.SetMapIndex(key, vsrc)
	}

	return nil
}
