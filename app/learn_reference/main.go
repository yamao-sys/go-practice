package main

import (
	"fmt"
)

// // 可変長引数
// func Sum(s ...int) int {
// 	n := 0
// 	for _, v := range s {
// 		n += v
// 	}
// 	return n
// }

// func receiver(c chan int) {
// 	for {
// 		i := <-c
// 		fmt.Println(i)
// 	}
// }

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	// var sl []int // 整数型のslice. sliceは要素数未指定
	// fmt.Println(sl)

	// var sl2 []int = []int{100, 200}
	// fmt.Println(sl2)

	// sl3 := []string{"A", "B"}
	// fmt.Println(sl3)

	// sl4 := make([]int, 5)
	// fmt.Println(sl4)

	// sl2[0] = 1000
	// fmt.Println(sl2)

	// sl5 := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl5)

	// fmt.Println(sl5[2:4])            // indexが2から4の手前まで(つまり2〜3)
	// fmt.Println(sl5[:2])             // index2の手前まで(つまり0~1)
	// fmt.Println(sl5[2:])             // index2から最後まで
	// fmt.Println(sl5[:])              // 全部
	// fmt.Println(sl5[len(sl5)-1])     // 最後の要素
	// fmt.Println(sl5[1 : len(sl5)-1]) // index1〜最後の手前まで(つまり1〜3)

	// sl1 := []int{100, 200}
	// fmt.Println(sl1)

	// sl1 = append(sl1, 300)
	// fmt.Println(sl1)

	// sl1 = append(sl1, 400, 500, 600)
	// fmt.Println(sl1)

	// sl21 := make([]int, 5)
	// fmt.Println(sl21)

	// fmt.Println(len(sl21))

	// fmt.Println(cap(sl21)) // sliceの容量

	// sl31 := make([]int, 5, 10) // 要素数は5, 容量は10

	// fmt.Println(len(sl31))

	// fmt.Println(cap(sl31)) // sliceの容量が10

	// sl41 := append(sl31, 1, 2, 3, 4, 5, 6, 7)
	// fmt.Println(len(sl41))
	// fmt.Println(cap(sl41)) // 容量を超えた要素数になると、容量は倍増していく

	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl) // 参照型(slice, map, channel)はメモリアドレスを代入するため、sl2とslは同じアドレスを指すことになる

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2) // 基本型は別の値

	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// fmt.Println(sl2)
	// n := copy(sl2, sl) // nにはcopyした要素数が入る
	// fmt.Println(n, sl2)

	// sl := []string{"A", "B", "C"}
	// fmt.Println(sl)

	// for i := range sl {
	// 	fmt.Println(i, sl[i])
	// }

	// for i := 0; i < len(sl); i++ {
	// 	fmt.Println(sl[i])
	// }

	// fmt.Println(Sum(1, 2, 3))

	// fmt.Println(Sum())

	// sl := []int{1, 2, 3}
	// fmt.Println(Sum(sl...))

	// var m = map[string]int{"A": 100, "B": 200}
	// fmt.Println(m)

	// m2 := map[string]int{"A": 100, "B": 200}
	// fmt.Println(m2)

	// m3 := map[int]string{
	// 	1: "A",
	// 	2: "B",
	// }
	// fmt.Println(m3)

	// m4 := make(map[int]string)
	// fmt.Println(m4)

	// m4[1] = "Japan"
	// m4[2] = "USA"
	// fmt.Println(m4)

	// fmt.Println(m["A"])
	// fmt.Println(m4[2])
	// fmt.Println(m4[3]) // 未登録のキーはデフォルト値(今回はstringなので空文字)

	// s, ok := m4[3] // 第二の返り値は取り出し成功かどうか
	// if !ok {
	// 	fmt.Println("error")
	// }
	// fmt.Println(s, ok)

	// m4[2] = "US"
	// fmt.Println(m4)

	// m4[3] = "CHINA"
	// fmt.Println(m4)

	// delete(m4, 3)
	// fmt.Println(m4)

	// fmt.Println(len(m4))

	// m := map[string]int{
	// 	"Apple":  100,
	// 	"Banana": 200,
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// for k := range m {
	// 	fmt.Println(k)
	// }

	// channel
	// 複数のゴルーチン間でデータの受け渡しをするために設計されたデータ構造
	// 宣言・操作
	// var ch1 chan int
	// // var ch2 <-chan int // 受信専用のchannel
	// // var c3 chan<- int // 送信専用のchannel

	// ch1 = make(chan int)

	// ch2 := make(chan int)

	// fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	// ch3 := make(chan int, 5) // キューのバッファサイズ5
	// fmt.Println(cap(ch3))

	// ch3 <- 1
	// fmt.Println(len(ch3))

	// ch3 <- 2
	// ch3 <- 3
	// fmt.Println(len(ch3))

	// i := <-ch3
	// fmt.Println(i)
	// i2 := <-ch3
	// fmt.Println(i2)
	// fmt.Println(len(ch3))

	// fmt.Println(<-ch3)
	// fmt.Println(len(ch3))

	// ch3 <- 1
	// fmt.Println(<-ch3)
	// ch3 <- 2
	// ch3 <- 3
	// ch3 <- 4
	// ch3 <- 5
	// ch3 <- 6
	// fmt.Println(len(ch3))

	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// // fmt.Println(<-ch1)

	// // receiverをch1, ch2の2つで並行して実行
	// // receiverではそれぞれのchのデータ受信を待つ
	// go receiver(ch1)
	// go receiver(ch2)

	// i := 0
	// for i < 100 {
	// 	// ch1, ch2それぞれにデータを送信 → receiverでそれぞれ受け取って処理
	// 	ch1 <- i
	// 	ch2 <- i
	// 	time.Sleep(50 * time.Millisecond)
	// 	i++
	// }

	// ch1 := make(chan int, 2)
	// ch1 <- 1
	// close(ch1)
	// // ch1 <- 1

	// // fmt.Println(<-ch1) // closeされていても受信はできる
	// i, ok := <-ch1
	// fmt.Println(i, ok) // okはchannelのバッファ0 かつ closeされていればfalse

	// i2, ok := <-ch1
	// fmt.Println(i2, ok) // okはchannelのバッファ0 かつ closeされていればfalse

	// ch1 := make(chan int, 2)
	// go receiver("1.goroutin", ch1)
	// go receiver("2.goroutin", ch1)
	// go receiver("3.goroutin", ch1)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	i++
	// }
	// close(ch1)
	// time.Sleep(3 * time.Second)

	// ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// close(ch1)
	// for i := range ch1 {
	// 	fmt.Println(i)
	// }
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"
	// ch1 <- 1
	// ch2 <- "B"
	// ch1 <- 2

	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1)
	// fmt.Println(v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// receiver
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("receive", i3)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100 {
		// 	break
		// }
	}
}
