package base

import (
	"fmt"
	"runtime"

)
const SIZE = 1 << 12

type StackError struct {
	Filename      string
	CallingMethod string
	Line          int
	ErrorMessage  string
	StackTrace    string
}

func New(err interface{}) *StackError {
	return New_Skip(err,1)
}
func New_Skip(err interface{},skip int) *StackError {
	var errMessage string
	switch t := err.(type) {
	case *StackError:
		return t
	case string:
		errMessage = t
	case error:
		errMessage = t.Error()
	default:
		errMessage = fmt.Sprintf("%v", t)
	}
	stackErr := &StackError{}
	stackErr.ErrorMessage = errMessage
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		stackErr.Line = line
		stackErr.Filename=file
	}

	buf := make([]byte, SIZE)
	n := runtime.Stack(buf, false)
	stackErr.StackTrace = string(buf[:n])
	return stackErr
}
func (this *StackError) Error() string {
	return this.ErrorMessage
}
func (this *StackError) Stack() string {
	return fmt.Sprintf("\n%s:%d %s\n\t[Stack Info]:\n %s", this.Filename, this.Line, this.ErrorMessage, this.StackTrace)
}
func (this *StackError) Detail() string {
	return fmt.Sprintf("\n%s:%d %s\n", this.Filename, this.Line, this.ErrorMessage)
}
