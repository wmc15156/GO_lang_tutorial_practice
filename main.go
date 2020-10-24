// main 폴더부터 찾아서 컴파일한다.

// go 언어는 특정 함수를 찾게 되는데 go 프로그램의 시작점이 되는 부분

// 컴파일러는  package main와 그 안에 있는 main function을 먼저 찾고 실행시킴

// 또 변수는 선언을 했으면 사용을 해야됨 // 변수명 declared but not used

package main

import (
	"fmt"
	"strings"
)

// test := "test입니다"

func multiply(a int, b int) int {
	// 파리미터 값, return을 타입을 정해야됨
	// a,b 둘다 int일 경우에는 a,b int 이렇게 해줘도 된다.
	return a * b
}

func practiceFunc(words ...string) {
	fmt.Println(words)
}

func practiceFunc2(word string) (int, string) {
	// return 값으로 여러개를 할 수 있다.
	return len(word), strings.ToUpper(word)
}

func practiceFunc3(word string) (length int, uppercase string) {
	// return 값으로 여러개를 할 수 있다.

	// return 타입 지정시 반환되는 변수명도 지정할 수 있다.
	// 마지막에 리턴할때 리턴변수를 명시하지 않아도 됨.
	length = len(word)
	uppercase = strings.ToUpper(word)

	return
}

func practiceFunc4() string {
	// defer 해당함수의 종료 직전(return전)에 실행된다.
	defer fmt.Println("저의 순서가 궁금합니다.")

	fmt.Println("저는 제일 첫번째로 출력될거 같아요")

	return "제가 제일 마지막이겠군요"
}

func main() {

	const coding string = "어렵다"

	// coding = "me too" // cannot assign to coding error 발생

	var lang string = "hello" // string 타입만 가능함

	lang = "false" // boolean 타입 선언 시 에러 발생

	fmt.Println(lang) //cannot use false (type untyped bool) as type string in assignment

	name := "hello" // 축약식 => func안에서만 선언을 해줘야지 가능함

	fmt.Println(name) // hello

	// fmt.Println(test) // syntax error: non-declaration statement outside function body

	fmt.Println(multiply(2, 4)) // 8

	practiceFunc("h", "e", "l", "l", "o")

	fmt.Println(practiceFunc2("James")) // 5 JAMES

	fmt.Println(practiceFunc3("go practice4")) // 12 GO PRACTICE4

	fmt.Println(practiceFunc4())

	// 출력
	// 저는 제일 첫번째로 출력될거 같아요
	// 저의 순서가 궁금합니다.
	// 제가 제일 마지막이겠군요.

}
