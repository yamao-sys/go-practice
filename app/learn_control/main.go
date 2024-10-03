package main

import (
	"fmt"
	"strconv"
	"time"
)

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

// deferは後から登録された式から評価される 出力順は3→2→1
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}

func main() {
	// if
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I' don't know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		// scopeのxはif文のもの
		fmt.Println(x)
	}
	// このscopeはif文外のもの
	fmt.Println(x)

	var s string = "A"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	// ite := 0
	// for {
	// 	ite++
	// 	if ite == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i = 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }
	// for i, _ := range arr {
	// 	fmt.Println(i)
	// }

	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apply": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	for k := range m {
		fmt.Println(k)
	}

	// n := 5
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I' don't know")
	// }

	// // switch文内のみのscopeのn
	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I' don't know")
	// }

	// n := 6
	// if n > 0 && n < 4 {
	// 	fmt.Println("0 < n < 4")
	// }
	// switch n {
	// case n > 0 && n < 4:
	// 	fmt.Println("0 < n < 4")
	// case n > 3 && n < 7:
	// 	fmt.Println("3 < n < 7")
	// }

	anything("aaa")
	anything(1)

	var x1 interface{} = 3
	xx, isInt := x1.(int) // intの型アサーション
	fmt.Println(xx+2, isInt)

	f, isFloat64 := x1.(float64)
	fmt.Println(f, isFloat64)

	// if文を使った型アサーション
	if x1 == nil {
		fmt.Println("None")
	} else if xx, isInt := x1.(int); isInt {
		fmt.Println(xx, "x1 is int")
	} else if s, isString := x1.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	// switch文を使った型アサーション
	switch x1.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x1.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}

	// ラベル付きfor
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理をしない")
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END")
	// continueと組み合わせる
	// Loop:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 3; j++ {
	// 			if j > 1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(j, j, i*j)
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}

	// defer 関数の終了時に行う処理
	TestDefer()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	RunDefer()

	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close() // deferによるリソースの解放

	// file.Write([]byte("Hello"))

	// panic と recover
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error")
	// fmt.Println("Start")

	// go goroutin 並行処理を簡単に
	// go sub()
	// go sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	// init
	fmt.Println("Main")
}
