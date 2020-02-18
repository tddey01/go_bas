package main

import "os"

/*
	golang 判断一个文件夹是否存在的方法使用os.Stat()函数返回的错误值进行判断
	1、如果返回错误为nil 说明文件或者文件存在
	2、如果返回错误雷士用os.IsNotExist()判断为true， 说明文件或者文件夹不存在
	3、 如果返回错误为其他类型，怎不确定是否存在


*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil { // 文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func main() {

}
