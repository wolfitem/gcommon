package base


import (

	"strings"
)

// 判断 list中是否包含a字符串
func StringInSlice(a string, list *[]string) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}



// 字符串切分转成array
func StringToSlice(text, separator string) (list []string) {
	return strings.Split(text, separator)
}

// 去除前后空间
func TrimSpace(s string) string {
    return strings.Trim(s, " ")
}
