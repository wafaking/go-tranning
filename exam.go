package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Reader interface {
	Prints()
}
type Writer interface {
	Gets()
}

type ReadWriter interface {
	Writer
	Reader
}

type Student struct {
	Name string
	Age  int
}

func (p *Student) Prints() {
	fmt.Printf("[name]=prints...\n")
}
func (p *Student) Gets() {
	fmt.Println("gets... ...")
}
func Test(f ReadWriter) {
	f.Prints()
	f.Gets()
}

func HstInterface(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case *Student:
			fmt.Println(111)
		case string:
			fmt.Println(222)
		case int32, int, int64:
			fmt.Println(333)
		}

	}

}

func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		fmt.Println("u2s, err: ", err)
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

func main_1() {
	// str := `sfds\u4e2d双方都\u56fd  {"a":"\u003cp\u003e中国\u003c/p\u003e","b":"hello"}`
	str := "\u624d\u827a\u6280\u80fd"
	buf := bytes.NewBuffer(nil)

	i, j := 0, len(str)
	fmt.Println(i, j)
	for i < j {
		x := i + 6
		if x > j {
			buf.WriteString(str[i:])
			break
		}
		if str[i] == '\\' && str[i+1] == 'u' {
			hex := str[i+2 : x]
			r, err := strconv.ParseUint(hex, 16, 64)
			if err == nil {
				buf.WriteRune(rune(r))
			} else {
				buf.WriteString(str[i:x])
			}
			i = x
		} else {
			buf.WriteByte(str[i])
			i++
		}
	}
	//sfds中双方都国  {"a":"中国","b":"hello"}
	fmt.Println(buf.String())
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func funct(num int) func(res int) int {
	return func(res int) int {
		res += num
		return res
	}
}
