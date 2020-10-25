package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// pointer 예제
	// & 해당 주소값을 볼수 있음
	// *는 해당 메모리주소를 살펴볼 수 있음(값을 확인 할 수 있다.)
	// * 이용해 주소에 담긴값을 변경 할 수 있음

	favFood := []string{"kimchi", "ramen"}

	nico := person{name: "kim", age: 18, favFood: favFood}

	fmt.Println(nico)

	a := 2
	b := &a // &로 a의 메모리주소를 b에 할당함

	fmt.Println(&a, b) // 0xc0000140a0 0xc0000140a0

	*b = 4 // *로 b변수 메모리에 4를 할당 함

	fmt.Println(a, *b) // 4 4 서로 같은 메모리를 참조하기때문에 똑같은값이 출력

	c := 6
	d := &c

	*d = 10

	fmt.Println(c, *d, &c, d) // 10 10 0xc0000b2020 0xc0000b2020

	// 배열

	names := [5]string{"k", "h", "Go", "Learn", "1"} // 자리수 및 타입을 지정해줘야됨
	fmt.Println(names)

	array := []string{"a", "r", "r", "a", "y"} // slice

	array = append(array, "test")
	fmt.Println(array[0])

	nicos := map[string]string{"name": "nico", "age": "12"}

	fmt.Println(nicos)
	for key, value := range nicos {
		fmt.Println(key, value)
	}
}
