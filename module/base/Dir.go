
package base


import (
	"os"
	"fmt"
)


//创建文件夹
func CreateDir(folder string) error {
	err := os.Mkdir(folder, 0755) // 第二个参数为文件夹的权限设置（这里使用默认值）
	if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("成功创建文件夹")
    }
	return err;
}
