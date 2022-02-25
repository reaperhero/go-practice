package main

import (
	"reflect"
	"testing"
)

type ServiceConfig struct {
	Service map[string]struct {
		Config map[string]interface{} `config:"config" json:",omitempty"`
	} `config:"service" validate:"required"`
}

func allElem(v reflect.Value) reflect.Value {
	for v.IsValid() && (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) {
		v = v.Elem()
	}
	return v
}

func TestRefectString(t *testing.T) {


}
