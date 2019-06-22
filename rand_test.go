package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	// rand.Seed(time.Now().Unix())
	// value := rand.Intn(10)
	now := time.Now().Unix()
	fmt.Println(now)
	rand.Seed(now)
	randNum := rand.Intn(5) + 3
	fmt.Println(randNum)
}
