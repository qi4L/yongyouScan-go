package monitorservlet_Desera

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
	url1 := url + "service/monitorservlet"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Post(url1)
	if err != nil {
	} else {
		if resp.Status == "200 OK" {
			color.Green.Println("[+] 用友 NC MonitorServlet反序列化漏洞存在 -> " + url)
			return
		}
	}
	url2 := url + "servlet/~ic/nc.bs.framework.mx.monitor.MonitorServlet"
	resp1, err1 := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Post(url2)
	if err1 != nil {
		color.Red.Println("[-] 用友 NC MonitorServlet反序列化漏洞不存在")
		return
	}
	if resp1.Status == "200 OK" {
		color.Green.Println("[+] 用友 NC MonitorServlet反序列化漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 NC MonitorServlet反序列化漏洞不存在")
}
