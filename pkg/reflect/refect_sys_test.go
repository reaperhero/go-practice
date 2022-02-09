package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type SysConfig struct {
	Platform *Platform `sys:"platformsecurity"`
}

type Platform struct {
	LoginEncrypt    string `sys:"login_encrypt"`
	LogoutEncrypt   int    `sys:"logout_encrypt"`
	LogSliceEncrypt []int  `sys:"loglice_encrypt"`
}

type searchKey struct {
	key string
}

func (s *searchKey) add(key string) string {
	if s.key == "" {
		s.key = key
	} else {
		s.key = s.key + "." + key
	}
	return s.key
}

func (s *searchKey) delTail() {
	if s.key == "" {
		return
	}
	arr := strings.Split(s.key, ".")
	s.key = strings.Join(arr[0:len(arr)-1], ".")
}

func TestSys(ttt *testing.T) {
	var (
		splitKey    = "sys"
		traverse    func(target interface{})
		combination = searchKey{}
		config      = SysConfig{
			Platform: &Platform{
				LoginEncrypt:  "11",
				LogoutEncrypt: 11,
			},
		}
		data = map[string]interface{}{
			"platformsecurity.login_encrypt":   "22",
			"platformsecurity.logout_encrypt":  "22",
			"platformsecurity.loglice_encrypt": []int{1, 2, 3},
		}
	)

	traverse = func(target interface{}) {
		sVal := reflect.ValueOf(target)
		sType := reflect.TypeOf(target)
		if sType.Kind() == reflect.Ptr {
			sVal = sVal.Elem()
			sType = sType.Elem()
		}

		num := sVal.NumField()
		for i := 0; i < num; i++ {
			k := sType.Field(i).Tag.Get(splitKey)
			if k != "" {
				combination.add(k)
			}
			//判断字段是否为结构体类型，或者是否为指向结构体的指针类型
			if sVal.Field(i).Kind() == reflect.Struct || (sVal.Field(i).Kind() == reflect.Ptr && sVal.Field(i).Elem().Kind() == reflect.Struct) {
				traverse(sVal.Field(i).Interface())
			} else {
				field := sVal.Field(i)
				v, ok := data[combination.key]
				if field.IsValid() && ok {
					switch field.Type().Kind() {
					case reflect.String:
						if v, ok := v.(string); ok {
							sVal.Field(i).Set(reflect.ValueOf(v))
						}
					case reflect.Int:
						if v, ok := v.(int); ok {
							field.Set(reflect.ValueOf(v))
						}
					case reflect.Slice:
						if field.Type().Elem().Kind() == reflect.Int {
							if op, ok := v.([]int); ok {
								field.Set(reflect.ValueOf(op))
							}

						}
						if field.Type().Elem().Kind() == reflect.String {
							if op, ok := v.([]string); ok {
								field.Set(reflect.ValueOf(op))
							}
						}
					}
				}
			}
			if k != "" {
				combination.delTail()
			}
		}
	}
	traverse(&config)
	fmt.Printf("%v", config.Platform)
}
