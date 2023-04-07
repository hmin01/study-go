package main

import (
	"fmt"
	"time"
)

func main() {
	// 문자열 채널 2개 생성
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 채널을 이용한 고루틴 1
	go func() {
		for {
			time.Sleep(5 * time.Second)
			ch1 <- "one"
		}
	}()
	// 채널을 이용한 고루틴 2
	go func() {
		for {
			time.Sleep(10 * time.Second)
			ch2 <- "two"
		}
	}()

	// 채널로 수신되는 메시지를 처리하는 메인 프로세스
	// for 문을 이용하여 반복적으로 진행되며,
	// select 문을 이용하여 case에 따라 각각의 작업이 수행됩니다.
	// default 가 없을 경우, case의 조건이 맞을 때까지 select 문은 블락(Block)됩니다.
	for {
		fmt.Println("//----- start -----//")
		select {
		case msg := <-ch1:
			fmt.Println("received: ", msg)
		case msg := <-ch2:
			fmt.Println("received: ", msg)
		default:
			fmt.Println("default")
		}
		fmt.Println("//----- end -----//")
	}
}
