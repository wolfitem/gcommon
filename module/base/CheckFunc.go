package base

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CheckFunc func(v interface{}, param string) (err error)

func is_empty(v interface{}) bool {
	return v == "" || is_nil(v)
}

func is_nil(v interface{}) bool {
	return v == nil
}

func is_positive(v int) bool {
	return v > 0
}

func is_true(v string) bool {
	return strings.ToLower(v) == "true"
}

func is_false(v string) bool {
	return strings.ToLower(v) == "false"
}

func IsTrue(v interface{}, param string) (err error) {
	if !is_true(fmt.Sprint(v)) {
		err = errors.New(param)
	}
	return
}

func IsFalse(v interface{}, param string) (err error) {
	if !is_false(fmt.Sprint(v)) {
		err = errors.New(param)
	}
	return
}

func IsPositive(n interface{}, param string) (err error) {
	number, err := strconv.ParseInt(fmt.Sprint(n), 10, 0)
	if err != nil {
		msg := fmt.Sprintf("%s is not number !", param)
		err = errors.New(msg)
	}
	if !is_positive(int(number)) {
		msg := fmt.Sprintf("%s is not positive number ! value = %d", param, number)
		err = errors.New(msg)
	}
	return
}

func IsNotEmpty(v interface{}, param string) (err error) {

	if is_empty(v) {
		msg := fmt.Sprintf("%s is empty!", param)
		err = errors.New(msg)
	}
	return
}

func IsNil(v interface{}, param string) (err error) {

	if !is_nil(v) {
		msg := fmt.Sprintf("%s is not nil!", param)
		err = errors.New(msg)
	}
	return
}

func IsNotNil(v interface{}, param string) (err error) {

	if is_nil(v) {
		msg := fmt.Sprintf("%s is nil!", param)
		err = errors.New(msg)
	}
	return
}

func IsEmpty(v interface{}, param string) (err error) {

	if !is_empty(v) {
		msg := fmt.Sprintf("%s is not empty!", param)
		err = errors.New(msg)
	}
	return
}

//todo 这个方法在实际应用中一般不可能传递 []interface{},需要重写
func HasAny(v []interface{}, param string) (err error) {

	if v == nil {
		msg := fmt.Sprintf("%s is null!", param)
		err = errors.New(msg)
	}
	if len(v) == 0 {
		msg := fmt.Sprintf("%s  count must > 0!", param)
		err = errors.New(msg)
	}
	return
}
