package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestTimeTicker 定时器学习
func TestTimeTimer(t *testing.T) {
	// time1 := time.NewTimer(-5 * time.Second)
	fmt.Printf("now: %v\n", time.Now())
	// expire := <-time1.C
	expire := <-time.NewTimer(3 * time.Second).C
	fmt.Printf("end: %v\n", expire)
	return

}

func TestSelectTicker(t *testing.T) {
	var (
		ticker *time.Ticker
		wg     sync.WaitGroup
	)
	ticker = time.NewTicker(1 * time.Second)
	wg.Add(5)

	go func() {
		for {
			<-ticker.C
			fmt.Println(1111)
			wg.Done()
		}
	}()
	wg.Wait()
	return
}

func TestTimer(t *testing.T) {
	time.AfterFunc(2*time.Second, func() {
		fmt.Printf("expired1\n")
	})
	fmt.Println(33)

	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Printf("expired2\n")

	chaner := time.After(2 * time.Second)
	<-chaner
	fmt.Printf("expired3\n")
}
