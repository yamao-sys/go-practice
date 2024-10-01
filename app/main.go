package main // package宣言は1ファイルに1つ. これをファイル間でimportして使用する

import (
	"fmt"
	"time"
)

// Hello world
/*
複数行の
コメント
*/

func main() {
	fmt.Println("hello world!!")
	fmt.Println(time.Now())
}
