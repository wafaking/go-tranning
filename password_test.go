package main

import (
	"fmt"
	"testing"

	"github.com/go-trellis/algorithm/encryption/hash"
)

func TestPasswd(t *testing.T) {
	salt := "8773711854"
	// passwd := "99henanquan"
	passwd := "amo123456"
	// passwd := "amo1234567"
	// fmt.Println(time.Unix(1552728402, 0))
	// // time.Unix(date, 0).Weekday() == time.Wednesday
	two := hash.NewMD5().SumTimes(passwd, 2)
	three := hash.NewMD5().SumTimes(passwd, 3)
	four := hash.NewMD5().Sum(three + salt)
	fmt.Println("two", two)
	fmt.Println("three", three)
	fmt.Println("four", four)

}
