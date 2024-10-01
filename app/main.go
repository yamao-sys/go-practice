package main // package宣言は1ファイルに1つ. これをファイル間でimportして使用する

import (
	"fmt"
)

// Hello world
/*
複数行の
コメント
*/

// i5 := 500 // 暗黙的な定義は関数内でのみ可能
var i5 int = 500 // 明示的な定義は関数外でも使用可能

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// fmt.Println("hello world!!")
	// fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 変数の宣言だけだと、型のデフォルト値が表示される

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 500 // 再定義はNG

	// i4 = "Hello" // 異なる型の代入NG
	fmt.Println(i5)

	outer()

	// fmt.Println(s4) // 外部関数内の変数は使えず. s4はouterのローカル変数
	// var s5 string = "Not used" // 定義した変数が未使用だとエラーになる
}
