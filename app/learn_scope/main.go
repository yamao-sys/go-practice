package main

import (
	"app/learn_scope/foo"
	. "fmt"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

// 返り値の変数名を指定
func Do(s string) (b string) {
	b = s
	// ブロック内のbはブロック内のscope
	{
		b := "BBBB"
		Print(b)
	}
	// ブロックを抜けると、このメソッド内のscope
	Print(b)
	return b
}

func main() {
	Println(foo.MAX)
	// fmt.Println(foo.min) // privateなので参照NG
	Println(foo.ReturnMin())

	Println(appName())
	Println(Do("A"))
}
