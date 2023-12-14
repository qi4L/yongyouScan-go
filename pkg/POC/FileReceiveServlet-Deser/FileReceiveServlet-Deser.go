package FileReceiveServlet_Deser

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

// https://github.com/Threekiii/Awesome-POC/blob/master/OA%E4%BA%A7%E5%93%81%E6%BC%8F%E6%B4%9E/%E7%94%A8%E5%8F%8B%20NC%20FileReceiveServlet%20%E5%8F%8D%E5%BA%8F%E5%88%97%E5%8C%96RCE%E6%BC%8F%E6%B4%9E.md

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/servlet/FileReceiveServlet"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 FileReceiveServlet反序列化漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 FileReceiveServlet反序列化漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 FileReceiveServlet反序列化漏洞不存在")
}
