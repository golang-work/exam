package main

import (
	"fmt"
	"path"
	"qing358/path/util"
)

func main() {
	fmt.Println(path.IsAbs("aa/bb/bb"))
	fmt.Println(path.IsAbs("/aa/bb/bb"))
	fmt.Println(path.Split("/aa/bb/bb"))
	fmt.Println(path.Dir("/aa/bb/bb"))
	fmt.Println(path.Base("/aa/bb/bb"))
	fmt.Println(path.Ext("/aa/bb/bb.txt"))
	fmt.Println(path.Join("/aa/bb/", "cc", "", "/ee", "//ff/"))
	fmt.Println(path.Clean("/a/b/c/../bb/bb.txt"))
	//fmt.Println(os.Getwd())
	fmt.Println(util.GetCurrentPath())
}
