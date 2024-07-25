package main

import (
	"fmt"
	"reflect"
	"strings"
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

func ArrayToString(source []int64, delim string) string {

	return strings.Trim(strings.Replace(fmt.Sprint(source), " ", delim, -1), "[]")

}
func TestRefectString(t *testing.T) {
	fmt.Println(ArrayToString([]int64{1, 3, 5}, ","))
}
