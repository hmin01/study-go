package main

import (
	"fmt"
	"time"
)

func main() {
	ExampleTicker()

	// 프로세스 대기 (5분)
	time.Sleep(5 * time.Minute)
	// 종료 알림
	fmt.Println("[NOTICE] Ticker stopped.")
}

func ExampleTicker() {
	// 시간 간격 생성
	var interval time.Duration = 1 * time.Second
	// 티커 생성
	ticker := time.NewTicker(interval)

	// 루프
	go TimeLoop(ticker)
}

func TimeLoop(ticker *time.Ticker) {
	for t := range ticker.C {
		fmt.Printf("\r%s", t.Format("2006-01-02 15:04:05"))
	}
	defer ticker.Stop()
}
