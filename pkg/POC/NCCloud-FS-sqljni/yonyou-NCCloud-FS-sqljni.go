package NCCloud_FS_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest()
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/fs"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url)
	if err != nil {
		color.Red.Println("[-] 用友 NCCloud FS 文件管理 SQL 注入不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 NCCloud FS 文件管理 SQL 注入存在，只访问了fs，记得sqlmap跑下username参数来确认 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 NCCloud FS 文件管理 SQL 注入不存在")
}
