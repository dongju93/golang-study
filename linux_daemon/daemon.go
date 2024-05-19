package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// handleHello는 "/hello" 경로로의 요청을 처리하는 핸들러 함수입니다.
func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// "/hello" 경로로 요청이 올 때 handleHello 함수가 실행됩니다.
	http.HandleFunc("/hello", handleHello)

	// 8080 포트에서 HTTP 서버를 시작합니다.
	go func() {
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// 시스템 신호를 처리하기 위해 채널을 설정합니다.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 신호가 수신될 때까지 대기합니다.
	<-quit
	log.Println("Shutting down server...")
}

// arm64 아키텍처 리눅스용 바이너리를 빌드합니다.
// GOARCH=arm64 GOOS=linux go build -o bin/hello_daemon linux_daemon/daemon.go
