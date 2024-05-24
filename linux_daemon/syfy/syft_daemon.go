package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// handleHello는 "/hello" 경로로의 요청을 처리하는 핸들러 함수입니다.
func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// handleSBOM은 "/sbom" 경로로의 요청을 처리하는 핸들러 함수로,
// syft 명령어를 실행하여 SBOM을 생성합니다.
func handleSBOM(w http.ResponseWriter, r *http.Request) {
	// POST 요청만 처리합니다.
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 동일한 경로에 있는 syft 바이너리 실행.
	cmd := exec.Command("syft", "dir:/", "-o", "cyclonedx-json")

	// Output capture
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
		return
	}

	// JSON 화 후 응답
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	// "/hello" 경로로 요청이 올 때 handleHello 함수가 실행됩니다.
	http.HandleFunc("/hello", handleHello)
	// "/sbom" 경로로 요청이 올 때 handleSBOM 함수가 실행됩니다.
	http.HandleFunc("/sbom", handleSBOM)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13, // 최소 TLS 버전을 1.3으로 설정
	}

	server := &http.Server{
		Addr:      ":8064",
		TLSConfig: tlsConfig,
	}

	// 8064 포트에서 HTTP 서버를 시작합니다.
	go func() {
		log.Println("Starting server on :8064")
		if err := server.ListenAndServeTLS("cert/cert.pem", "cert/private.key"); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// 시스템 신호를 처리하기 위해 채널을 설정합니다.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 신호가 수신될 때까지 대기합니다.
	<-quit
	log.Println("Shutting down server...")
	if err := server.Close(); err != nil {
		log.Fatalf("Server Close: %v", err)
	}

}
