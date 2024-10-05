package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i, i2)

	fmt.Println(i + 50)
	// fmt.Println(i + i2) // 異なるint型の演算はNG

	fmt.Printf("%T\n", i2) // 型を調べる

	fmt.Printf("%T\n", int32(i2)) // 型変換

	fmt.Println(i + int(i2))

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2 // 暗黙的な定義はfloat64になる
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32) // 実際に使用するのはfloat64

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan)

	var t, f bool = true, false
	fmt.Println(t, f)

	var s string = "Hello Golang!!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
		test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0])         // byte型
	fmt.Println(string(s[0])) // 文字列はbyte配列の集まり

	byteA := []byte{72, 73} // slice
	fmt.Println(byteA)

	fmt.Println(string(byteA)) // 文字列はASCIIコード

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

	// 配列型は要素数を後から変更することはNG. 要素数を変更したい場合はslice型を使う
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"} // 文字列型の初期値は空文字なので、3つ目の要素は空文字
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1]) // 最後の要素は要素数 - 1

	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1 // 同じint型でも要素数が異なると、別の型と見なされる

	fmt.Println(len(arr1))

	var x interface{} // 全ての型を汎用的に表すinterface
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	x = 2
	// fmt.Println(x + 3) // interfaceとint型の演算はNG

	var i1 int = 1
	fl641 := float64(i1)
	fmt.Println(fl641)
	fmt.Printf("i1 = %T\n", i1)
	fmt.Printf("fl641 = %T\n", fl641)

	i21 := int(fl641)
	fmt.Printf("i2 = %T\n", i21)

	fl1 := 10.5
	i31 := int(fl1)
	fmt.Printf("i31 = %T\n", i31)
	fmt.Println(i31) // 小数をint型にすると、小数点以下切り捨て

	var s1 string = "100"
	fmt.Printf("s1 = %T\n", s1)

	i, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i22 int = 200
	s22 := strconv.Itoa(i22)
	fmt.Println(s22)
	fmt.Printf("s22 = %T\n", s22)

	// 文字列とbyte配列の変換
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
