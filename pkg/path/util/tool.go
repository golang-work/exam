package util

import "os"

func GetCurrentPath() string  {
	//path, err := os.Getwd()
	//if err != nil {
	//	p1, _ := os.Executable()
	//	return p1
	//}
	//
	//return path

	//p1, _ := os.Executable()
	p1, _ := os.Getwd()
	return p1
}
