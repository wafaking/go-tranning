package main

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
	var num = 3
	for i := 0; i < num+3; i++ {
		if i > num {
			goto label
		}else{}
	label:
		fmt.Println("1111")
	test:
		fmt.Println("2222")
	}
}
