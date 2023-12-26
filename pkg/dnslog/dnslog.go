package dnslog

import (
	"fmt"
	"github.com/imroc/req/v3"
	"math/rand"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest()
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func randCreator() string {
	str := "0123456789abcdefghigklmnopqrstuvwxyz"
	strList := []byte(str)
	result := []byte{}
	i := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < 26 {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
	return string(result)
}

func GetDnslogUrl() (string, string) {
	session := randCreator()
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
		"Cookie":     "PHPSESSID=" + session,
	}).Get("http://www.dnslog.cn/getdomain.php")
	if err != nil {
		if resp == nil {
			fmt.Println("与dns平台网络不可达,请检查网络")
			return "", ""
		}
	}
	return resp.String(), session
}

func GetDnslogRecord(PHPSESSID string) bool {
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
		"Cookie":     "PHPSESSID=" + PHPSESSID,
	}).Get("http://www.dnslog.cn/getrecords.php")
	if err != nil {
		fmt.Println("与dns平台网络不可达,请检查网络")
		return false
	}
	if resp.String() == "[]" {
		return true
	} else {

	}
	return false
}
