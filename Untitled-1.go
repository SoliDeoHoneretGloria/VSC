/*package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for i := 0; i < 4; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)

		}
	}()

	for i := 0; i < 4; i++ {
		<-done

		fmt.Println("메인 함수 : ", i)

		time.Sleep(time.Second)
	}
}
*/