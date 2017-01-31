package base

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

type CheckFunc func(v interface{},param string) (err error)

func is_empty(v interface{}) bool  {
	return  v == "" || is_nil(v)
}

func is_nil(v interface{}) bool  {
	return  v == nil
}

func is_positive(v int) bool{
	return v>0
}

func is_true(v string) bool {
	return strings.ToLower(v) == "true"
}

func is_false(v string) bool {
	return strings.ToLower(v) == "false"
}


func IsTrue(v interface{}, param string) (err error) {
	if !is_true(fmt.Sprint(v)) {
		err=errors.New(param)
	}
	return
}

func IsFalse(v interface{}, param string) (err error) {
	if !is_false(fmt.Sprint(v)) {
		err=errors.New(param)
	}
	return
}

func IsPositive(n interface{},param string) (err error) {
	number,err:=strconv.ParseInt(fmt.Sprint(n),10,0)
	if err!=nil{
		msg:=param + " is not number !"
		err=errors.New(msg)
	}
	if !is_positive(int(number)) {
		msg:=fmt.Sprintf("%s is not positive number ! value = %d" ,param,number)
		err=errors.New(msg)
	}
	return
}

func IsNotEmpty(v interface{}, param string) (err error)  {

	if is_empty(v){
		msg:=param + " is empty!"
		err=errors.New(msg)
	}
	return
}



func IsNil(v interface{}, param string) (err error)  {

	if !is_nil(v){
		msg:=param + " is not nil!"
		err=errors.New(msg)
	}
	return
}

func IsNotNil(v interface{}, param string) (err error) {

	if is_nil(v){
		msg:=param + " is nil!"
		err=errors.New(msg)
	}
	return
}

func IsEmpty(v interface{}, param string) (err error) {

	if !is_empty(v) {
		msg:=param + " is not empty!"
		err=errors.New(msg)
	}
	return
}

func HasAny(v []interface{}, param string) (err error)  {

	if v == nil {
		msg:=param + " is null!"
		err=errors.New(msg)
	}
	if len(v) == 0 {
		msg:=param + " count must > 0!"
		err=errors.New(msg)
	}
	return
}