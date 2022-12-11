package lookupcfg

import (
	"errors"
	"strconv"
	"strings"
)

// todo:
// uint8
// uint16
// uint32
// uint64
// int
// int8
// int16
// int32
// float32
// float64
// complex64
// complex128
// array
// map
// slice
// (are they really needed?)

// regards: https://github.com/fscdev/betterconf/blob/717ad842e676b299112b1278e5f93d0933dc71ac/betterconf/caster.py#L38
var booleanTrue []string = []string{"true", "on", "enable", "1", "yes", "ok"}
var booleanFalse []string = []string{"false", "off", "disable", "0", "no"}

func parseBool(x string) (bool, error) {
	x = strings.ToLower(x)
	for _, trueString := range booleanTrue {
		if x == trueString {
			return true, nil
		}
	}
	for _, falseString := range booleanFalse {
		if x == falseString {
			return false, nil
		}
	}
	return false, errors.New("the string can't be represented as boolean value")
}

func parseInt64(x string, bitSize int) (int64, error) {
	return strconv.ParseInt(x, 10, bitSize)
}

func parseFloat(x string, bitSize int) (float64, error) {
	return strconv.ParseFloat(x, bitSize)
}
