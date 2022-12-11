package lookupcfg

import (
	"fmt"
	"reflect"
	"strconv"
)

func panicf(format string, v ...any) {
	panic(fmt.Sprintf(format, v...))
}

func universalSet(fieldTypeKind reflect.Kind, fieldValue reflect.Value, value string) error {
	switch fieldTypeKind {
	case reflect.String:
		fieldValue.SetString(value)
		return nil

	case reflect.Int:
		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		fieldValue.SetInt(n)
		return nil

	}
	return nil
}
