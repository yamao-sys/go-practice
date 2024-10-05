package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(sl []int) {
	for i, v := range sl {
		sl[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	Double(n) // 基本型は別アドレスにコピーされて使用される → 元のnは不変
	fmt.Println(n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p) // 実態の値は*で表す(デリファレンス)

	// *p = 300
	// fmt.Println(n)

	// n = 200
	// fmt.Println(*p)

	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	Doublev3(sl) // 参照型は参照渡し
	fmt.Println(sl)
}
