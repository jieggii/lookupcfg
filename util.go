package lookupcfg

import (
	"errors"
	"fmt"
	"reflect"
)

func panicf(format string, v ...any) {
	panic(fmt.Sprintf(format, v...))
}

func universalSet(fieldTypeKind reflect.Kind, fieldValue reflect.Value, value string) error {
	switch fieldTypeKind {
	case reflect.Bool:
		result, err := parseBool(value)
		if err != nil {
			return err
		}
		fieldValue.SetBool(result)
	case reflect.Int:
		result, err := parseInt64(value, 64)
		if err != nil {
			return err
		}
		fieldValue.SetInt(result)
	case reflect.Float32:
		result, err := parseFloat(value, 32)
		if err != nil {
			return err
		}
		fieldValue.SetFloat(result)
	case reflect.Float64:
		result, err := parseFloat(value, 64)
		if err != nil {
			return err
		}
		fieldValue.SetFloat(result)
	case reflect.String:
		fieldValue.SetString(value)
	default:
		return errors.New(
			"this type is not supported (maybe yet) by lookupcfg. You can help me with that at https://github.com/jieggii/lookupcfg",
		)
	}
	return nil
}
