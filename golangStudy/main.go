package main

import (
	"awesomeProject/something"
	"fmt"
	"strings"
)

//함수 밖에서 이렇게
var name2 string = "nico"

//struct --> replace object in java
// go does not have construct method
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {

	//struct
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico.name)

	fmt.Println("hihi")
	// 함수이름 대문자로 해야 public export 가능
	something.SayBye()
	something.SayHello()

	const name string = "nico"
	const xx bool = false

	//함수안에서만 가능함, 변수 생성의 의미
	n := "nico"
	n = "lynn"
	fmt.Println(n)
	fmt.Println(multiply(2, 2))
	fmt.Println(lenAndUpper("david"))
	totalLength, _ := lenAndUpper("nicl")
	fmt.Println(totalLength)
	fmt.Println(lenAndUpper2("helen"))
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))

	// pointer
	a := 2
	b := &a
	fmt.Println(a, *b)

	//array
	names := [5]string{"nico", "lynn", "dal"}
	//slice
	nameslice := []string{"nico", "lynn", "dal"}
	nameslice[2] = "hehe"
	names[3] = "lala"

	// 하나를 추가하여 새로운 슬라이스 리턴
	nameslice = append(nameslice, "flynn")
	fmt.Println(names)

	//map
	nicoco := map[string]string{"name": "nico"}
	fmt.Println(nicoco)

	for key, value := range nicoco {
		fmt.Println(key, value)
	}

}

// 함수 정의 방식
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, upper string) {
	//defer는 함수가 리턴한후에 실행된다
	defer fmt.Println("I'm done")
	length = len(name)
	upper = strings.ToUpper(name)

	return
}

//for loop 형태
func superAdd(numbers ...int) (result int) {

	for index, number := range numbers {
		fmt.Println(index, number)
	}
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for _, number := range numbers {
		result += number
	}
	return
}

func canIDrink(age int) bool {
	// 조건 전에 변수 선언 가능
	if koreaAge := age + 2; koreaAge < 18 {
		return false
	} else {
		return true
	}
}
