// mygo project main.go
package main

import (
	"fmt"
)

var zkx2, zkx3 int = 1, 2

func main() {
	//	fmt.Println("Hello World!")
	//	//for 2 > 1 {
	//	fmt.Println("zkx ok")
	//	//}
	//	const a, b, c = 1, false, "str"
	//	println(a, b, c)

	//	var d int
	//	d = 1
	//	e := 2

	//	println(d)
	//	println(e)
	//	const (
	//		a = iota //0
	//		b        //1
	//		c        //2
	//		d = "ha" //独立值，iota += 1
	//		e        //"ha"   iota += 1
	//		f = 100  //iota +=1
	//		g        //100  iota +=1
	//		h = iota //7,恢复计数
	//		i        //8
	//	)
	//	fmt.Println(a, b, c, d, e, f, g, h, i)

	//	var zkx int = -1

	//	fmt.Println(&zkx)
	//	fmt.Println(*(&zkx))
	//	zkx = 2
	//	fmt.Println(&zkx)
	//	fmt.Println(*(&zkx))

	//	var p_zkx *int
	//	p_zkx = &zkx

	//	zkx = 3
	//	fmt.Println(*p_zkx)

	//	var b int = 15
	//	var a int

	//	//numbers := [6]int{1, 2, 3, 5}

	//	/* for 循环 */
	//	for a := 0; a < 10; a++ {
	//		fmt.Printf("a 的值为: %d\n", a)
	//	}

	//	for a < b {
	//		a++
	//		fmt.Printf("a 的值为: %d\n", a)
	//	}
	var a int = 10
	for a < 20 {
		if a == 15 {
			goto abc
		}
		a++
	}

abc:
	{
		fmt.Printf("我最喜欢的数字%d", a)
		a = 1
	}
	fmt.Printf("我最喜欢的数字%d", a)
}
