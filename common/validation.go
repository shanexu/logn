package common

import (
	"fmt"
	"github.com/elastic/go-ucfg"
	"reflect"
	"strings"
)

func init() {
	if err := ucfg.RegisterValidator("logn.oneof", func(v interface{}, params string) error {
		if v == nil {
			return nil
		}
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intSet, err := MakeIntSetFromStrings(strings.Split(params, " ")...)
			if err != nil {
				return err
			}
			if intSet.Has(int(val.Int())) {
				return nil
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintSet, err := MakeUintSetFromStrings(strings.Split(params, " ")...)
			if err != nil {
				return err
			}
			if uintSet.Has(uint(val.Uint())) {
				return nil
			}
		case reflect.Float32, reflect.Float64:
			floatSet, err := MakeFloatSetFromStrings(strings.Split(params, " ")...)
			if err != nil {
				return err
			}
			if floatSet.Has(val.Float()) {
				return nil
			}
		case reflect.String:
			if MakeStringSet(strings.Split(params, " ")...).Has(val.String()) {
				return nil
			}
		}
		return fmt.Errorf("requires value one of %q", params)
	}); err != nil {
		fmt.Println(err)
	}
}
