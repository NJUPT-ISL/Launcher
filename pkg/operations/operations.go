package operations

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func PostOperation(user string, pass string) (res bool) {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	data := url.Values{}
	data.Add("DDDDD",user)
	data.Add("upass",pass)
	data.Add("0MKKey","(unable to decode value)")
	data.Add("v6ip","")
	req, err := http.NewRequest("POST", "http://192.168.168.168/0.htm", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	if  strings.Contains(content, "msga=") {
		errReg, err := regexp.Compile("msga=([\\s\\S]*?);")
		if err != nil{
			fmt.Println("登录出现问题: ")
			fmt.Println(err)
			return false
		}
		if strings.Split(errReg.FindString(content),"'")[1] == ""{
			fmt.Println("登录出现认证问题: 您的账号可能已经登录。")
			return false
		}
		fmt.Println("登录出现认证问题: "+strings.Split(errReg.FindString(content),"'")[1])
		return false
	}
	return true
}

func GetOperation() {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://192.168.168.168/F.htm", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}