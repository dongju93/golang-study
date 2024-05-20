package main

import (
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

	// Create the command without output redirection
	cmd := exec.Command("syft", "dir:/", "-o", "cyclonedx-json")

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the output to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// Open the output file
// outputFile, err := os.Create("syclonedx_sbom.json")
// if err != nil {
// 	http.Error(w, fmt.Sprintf("Error creating output file: %v", err), http.StatusInternalServerError)
// 	return
// }
// defer outputFile.Close()

// Redirect the output to the file
// cmd.Stdout = outputFile
// cmd.Stderr = outputFile

// Run the command
// 	err = cmd.Run()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintln(w, "SBOM generation command executed successfully.")
// }

func main() {
	// "/hello" 경로로 요청이 올 때 handleHello 함수가 실행됩니다.
	http.HandleFunc("/hello", handleHello)
	// "/sbom" 경로로 요청이 올 때 handleSBOM 함수가 실행됩니다.
	http.HandleFunc("/sbom", handleSBOM)

	// 8080 포트에서 HTTP 서버를 시작합니다.
	go func() {
		log.Println("Starting server on :8064")
		if err := http.ListenAndServe(":8064", nil); err != nil {
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
