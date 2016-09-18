// mygo project main.go
package main

import (
	"fmt"
	"math"
)

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func swap(a, b int) (int, int) {
	return b, a
}

func bibao() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Student struct {
	name    string
	age     int
	address string
}

func (s Student) getStudentName() string {
	return s.name
}

func (s Student) getStudent() (string, int, string) {
	return s.name, s.age, s.address
}

func main() {
	var a int = max(3, 2)
	fmt.Println("max num:", a)
	fmt.Println(swap(100, 200))
	fmt.Println(math.Max(1.2, 3.4))
	fmt.Println(math.Sqrt(81))

	var x = bibao()
	fmt.Println(x())
	fmt.Println(x())
	x = bibao()
	fmt.Println(x())

	var s Student
	s.name = "zkx"
	s.age = 20
	s.address = "杭州"
	fmt.Println(s.getStudentName())
	fmt.Println(s.getStudent())

	var b = [5]int{}
	b[4] = 12
	for i := 0; i < 5; i++ {
		fmt.Println("::==", b[i])
	}

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	if ptr == nil {
		fmt.Println("hello")
	} else {
		fmt.Println("hi")
	}

	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country, city := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country], "::", city)
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	var ii int = 15
	fmt.Printf("%d 的阶乘是 %d\n", ii, Factorial(ii))
	fmt.Printf("%d\t", fibonaci(20))

}

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return result
}

func fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibonaci(n-2) + fibonaci(n-1)
}
