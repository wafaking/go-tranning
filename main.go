package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rest, _ := u2s(testStr)
	now := time.Now().Unix()
	fmt.Println(now)
	rand.Seed(now)
	randNum := rand.Intn(5) + 3
	fmt.Println(randNum)

	go func(num int) {
		ticker := time.NewTimer(4 * time.Second)
		for _, star := range stars {
			fmt.Println(num)
			<-ticker.C
		}
	}(num)
}
