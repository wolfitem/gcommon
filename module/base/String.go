package base


import (
	"strconv"
	"strings"
)

// 判断 list中是否包含a字符串
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//这个函数的第二个参数是一个函数
//这个函数将一个字符串转换为Int类型，如果失败了，则返回0，并输出错误。
func StringToInt(s string) (int64, error) {
	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
		return 0, err
	} else {
		return value, nil
	}
}

// 字符串切分转成array
func StringToSlice(text, separator string) (list []string) {
	return strings.Split(text, separator)
}
