package main

import "fmt"

const Pi = 3.14 // Goでは頭文字を大文字にすると、外部パッケージから呼び出せるようになる(パブリックになる)

const (
	URL      = "http://example.com"
	SiteName = "test"
)

// B, Cには1 E, Fには"A"が入る
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 連続する整数の連番を作成するiota
const (
	c0 = iota
	c1
	c2
)

// overflowで定義できず
// var Big int = 9223372036854775807 + 1
// 定数では可能
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)

	// Pi = 3 定数の上書きはNG

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}
