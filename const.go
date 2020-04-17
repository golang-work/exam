package main

import "fmt"

// 常量定义可以显示类型
const PI float64  = 3.1415926

//也可以由系统根据值来推导
const PI2 = 3.1415926

// 批量定义
const (
	DB_HOST = "127.0.0.1"
	DB_POST int32 = 3306
	DB_USERNAME = "root"
	DB_PASSWORD = "123456"
)

const Monday, Tuesday = 1, 2

// 枚举常量定义，当再次遇到iota时值会自动+1
//const (
//	READY int = iota
//	ALREADY = iota
//	PART = iota
//)

// 也可以这样写，不用多次写iota
const (
	READY = iota
	ALREADY
	PART
)

// 每遇到一次const关键字，iota值就会重置为0
const (
	STATUS_READY = iota
	STATUS_ALREADY
	STATUS_PART
)

func main() {
	fmt.Printf("PI value is: %f", PI)
	println()
	fmt.Printf("db host:%s, db post:%d, db username:%s, db password:%s", DB_HOST, DB_POST, DB_USERNAME, DB_PASSWORD)
	println()
	fmt.Printf("未开票:%d, 已开票:%d, 部分开票:%d", READY, ALREADY, PART)
	println()
	fmt.Printf("星期一:%d, 星期二:%d", Monday, Tuesday)
	println()
	fmt.Printf("待审核状态:%d, 已审核状态:%d, 部分审核状态:%d", STATUS_READY, STATUS_ALREADY, STATUS_PART)
}