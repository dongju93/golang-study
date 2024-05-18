package main

import (
	"fmt"     // 표준 입출력을 위한 패키지
	"os"      // 운영체제 기능을 위한 패키지, 여기서는 명령줄 인자를 처리하기 위해 사용
	"strconv" // 문자열을 숫자로 변환하기 위한 패키지
)

func main() {
	// 명령줄 인자가 4개인지 확인 (프로그램 이름 포함)
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./main <number> <operator> <number>")
		return // 인자 개수가 맞지 않으면 프로그램 종료
	}

	// 첫 번째 숫자 인자를 float64로 변환
	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	// 연산자 인자를 문자열로 저장
	operator := os.Args[2]
	// 두 번째 숫자 인자를 float64로 변환
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	// 변환에 실패하면 에러 메시지를 출력하고 종료
	if err1 != nil || err2 != nil {
		fmt.Println("Please provide valid numbers.")
		return
	}

	var result float64 // 연산 결과를 저장할 변수

	// 연산자에 따라 적절한 연산을 수행
	switch operator {
	case "+":
		result = num1 + num2 // 덧셈
	case "-":
		result = num1 - num2 // 뺄셈
	case "x":
		result = num1 * num2 // 곱셈
	case "/":
		// 나누는 수가 0이면 에러 메시지를 출력하고 종료
		if num2 == 0 {
			fmt.Println("Division by zero is not allowed.")
			return
		}
		result = num1 / num2 // 나눗셈
	default:
		// 유효하지 않은 연산자일 경우 에러 메시지를 출력하고 종료
		fmt.Println("Invalid operator. Use one of +, -, x, /.")
		return
	}

	// 결과가 정수인지 실수인지 확인하여 출력 형식을 다르게 함
	if result == float64(int(result)) {
		// 결과가 정수라면 정수로 출력
		fmt.Printf("Result: %d\n", int(result))
	} else {
		// 결과가 실수라면 불필요한 0을 제외하고 출력
		fmt.Printf("Result: %g\n", result)
	}
}
