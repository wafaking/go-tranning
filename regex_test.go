package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/astaxie/beego/logs"
)

var hansList = []string{
	"表示按unicode(utf-8)匹配",
	"中国，",
	"❤️会说话的刘二豆❤️",
	"可爱的朋哥哥?（朋家）",
	"小石头‼️承蒙厚爱",
	"姐姐Rachel和弟弟Ryan",
	"王小❗❗❗",
	"）&）&",
	"慧儿0-0",
	"华子哥♬",
}

func TestRegxHans(t *testing.T) {
	var compile = `[A-Za-z0-9\p{Han}]+`
	for i, hans := range hansList {
		match, err := regexp.MatchString(compile, hans)
		if err != nil {
			fmt.Println("regexp err: ", i, match, err)
		}
		fmt.Println("regexp res: ", i, match, err)
	}
}

func TestRegxEmoj(t *testing.T) {
	var compile = `[^\\u0000-\\uFFFF]`
	for i, text := range hansList {
		reg := regexp.MustCompile(compile)
		resList := reg.FindAllString(text, -1) // [bc ca olang]

		fmt.Println("regexp res: ", i, resList)
	}
}

func TestURL(t *testing.T) {
	var stringList = []string{
		"https://www.douyin.com/share/user/58558645805/",
		"https://www.amemv.com/share/user/93936144857?share_type=link",
		"http://www.iesdouyin.com/share/user/103509437955?u_code=17950dm0k&utm_campaign=client_share&app=aweme&utm_medium=ios&tt_from=cop",
		"https://live.kuaishou.com/profile/3xy2pxkk4xbpz4e",
		"https://live.kuaishou.com/profile/gaodi666",
		"testurl",
	}
	var compile = `^[\.\:a-zA-Z/]`
	for i, str := range stringList {
		// list := strings.Split(str, "/")
		// for _, v := range list {
		// res, err := regexp.Compile(compile)
		// if err != nil {
		// 	panic(err)
		// }
		// rest := res.ReplaceAllString(str, "")

		newURL := strings.Split(str, "?")[0]

		// }

		// lindex := strings.LastIndex(str, "/")
		// findex := strings.Index(str, "?")
		// if lindex >= 0 && findex >= 0 {
		// 	logs.Info("res1,  ", i, str[lindex:findex])
		// } else if lindex >= 0 && findex < 0 {
		// 	logs.Info("res2: ", i, str[lindex:])
		// } else {
		// 	logs.Info("res2: ", i, str)
		// }
		logs.Info("-----------")
	}

}
