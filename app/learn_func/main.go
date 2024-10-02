package main

import "fmt"

//	func Plus(x int, y int) int {
//		return x + y
//	}
func Plus(x, y int) int {
	return x + y
}

// 複数の戻り値
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 戻り値の変数名を指定
func Double(price int) (result int) {
	result = price * 2
	return
}

// 戻り値なしの関数
func Noreturn() {
	fmt.Println("No Return")
	return
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// 関数を引数に取る関数
func CallFunction(f func()) {
	f()
}

// クロージャの実装
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレータの実装: クロージャを活用して実装できる
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	i21, _ := Div(9, 3)
	fmt.Println(i21)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i1 := f(1, 2)
	fmt.Println(i1)

	i22 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i22)

	f1 := ReturnFunc()
	f1()

	CallFunction(func() {
		fmt.Println("I'm a function")
	})

	f2 := Later()
	fmt.Println(f2("Hello"))
	fmt.Println(f2("my"))
	fmt.Println(f2("name"))
	fmt.Println(f2("is"))
	fmt.Println(f2("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}
