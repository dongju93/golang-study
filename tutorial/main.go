// 실행 가능한 애플리케이션은 반드시 main 패키지를 포함
package main

// 표준 출력 및 입력을 위한 패키지
import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var x int

	x = 10

	// 변수를 선언하고 초기화
	y := 20

	k := add(x, y)

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// while 미지원, 무조건 for 문 사용
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 10 is:", sum)

	// hello world 출력
	fmt.Println("Hello, World!", x, y, k, sum)
}
