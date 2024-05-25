package main

import (
	// formatting 패키지
	"fmt"

	"nomadcoders.co/golang/scrapper/say"
)

func main() {
	// export를 위해선 대문자로 시작해야함
	fmt.Println("Hello, World!")
	say.SayHello()
	// Lowercase로 시작하면 export되지 않아서 다른 패키지에서 사용 불가
	// say.sayBye()
}
