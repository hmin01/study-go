package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 플래그
	urlPtr := flag.String("url", "", "a target url")
	// urlsPtr := flag.String("urls", "", "target urls")
	// 플래그 검증
	flag.Parse()
	// 예외 처리
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	// GET 호출
	res, err := http.Get(*urlPtr)
	// 예외 처리
	if err != nil {
		log.Fatalf("[ERROR] HTTP request error, %v\n", err)
	}
	defer res.Body.Close()

	target := res.Header.Get("date")
	targetTime, err := time.Parse(time.RFC1123, target)
	// 예외 처리
	if err != nil {
		log.Fatalf("[ERROR] Datetime parse error, %v\n", err)
	}

	// fmt.Println("Target: ", *urlPtr)
	// fmt.Println("=== Time diff ===")
	// fmt.Println("0: ", transformToDate(targetTime.Add(9*time.Hour), nil))
	// fmt.Println("1: ", transformToDate(time.Now(), nil))

	ticker(targetTime.Add(9 * time.Hour))
}

func transformToDate(datetime time.Time, format interface{}) string {
	if format != nil {
		return datetime.Format(format.(string))
	} else {
		return datetime.Format("2006-01-02 15:04:05")
	}
}

func ticker(target time.Time) {
	// 인터벌 정의
	var interval time.Duration = 1 * time.Second
	// 현재 시간
	targetTime := target
	hostTime := time.Now()
	// 티커 생성
	t := time.NewTicker(1 * time.Second)
	// Escape
	defer t.Stop()
	// Loop
	for {
		select {
		case <-t.C:
			// 시간 업데이트
			targetTime = targetTime.Add(interval)
			hostTime = hostTime.Add(interval)
			// 출력
			fmt.Printf("\r%s | %s", targetTime.Format("2006-01-02 15:04:05"), hostTime.Format("2006-01-02 15:04:05"))
		}
	}
}
