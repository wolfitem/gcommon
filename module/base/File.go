package base


import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

/**
 * 写文件
 */
func ReWriteFile(path string, context string) error {

	return ioutil.WriteFile(path, []byte(context), 0)

}

func WriteFile(path string, context string) error {
	var file *os.File
	var err error
	if CheckFileIsExist(path) {
		file, err = os.OpenFile(path, os.O_APPEND, 0666) //打开文件
		ErrorCheck(err)
	} else {
		file, err = os.Create(path)
		ErrorCheck(err)
	}


	_, err = file.WriteString(context)
	defer file.Close()
	ErrorCheck(err)
	return err
}

/**
 * 读文件
 */
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	fd, err := ioutil.ReadAll(fi)
	defer fi.Close()
	return string(fd)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//获取当前应用程序所在目录
func GetCurrentDirectory() string {
	file, _ := exec.LookPath(os.Args[0])

	path := filepath.Dir(file)
	return path
}

//获取目录下的文件夹及文件集合
func GetFilsAndDirs(path string) (files, dirs []string) {

	filepath.Walk(path,
		func(sub_path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				dirs = append(dirs, sub_path)
			} else {
				files = append(files, sub_path)
			}
			return nil
		})
	return
}