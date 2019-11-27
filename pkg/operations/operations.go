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

func DefaultLogin(user string, pass string) (res bool) {
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

func LoginChinaNetWifi(user string, pass string,ip string)  (res bool){
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
	data.Add("DDDDD",",0,"+user+"@njxy")
	data.Add("upass",pass)
	data.Add("R1","0")
	data.Add("R2","0")
	data.Add("R3","0")
	data.Add("R6","0")
	data.Add("para","00")
	data.Add("0MKKey","123456")
	data.Add("buttonClicked","")
	data.Add("redirect_url","")
	data.Add("err_flag","")
	data.Add("username","")
	data.Add("password","")
	data.Add("user","")
	data.Add("cmd","")
	data.Add("Login","")
	data.Add("v6ip","")
	req, err := http.NewRequest("POST", "http://p.njupt.edu.cn:801/eportal/?c=ACSetting&a=Login&protocol=http:&hostname=p.njupt.edu.cn&iTermType=1&wlanuserip="+ip+"&wlanacip=null&wlanacname=SPL_ME60&mac=00-00-00-00-00-00&ip="+ip+"&enAdvert=0&queryACIP=0&loginMethod=1", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func GetIP() string {
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
	req, err := http.NewRequest("GET", "http://p.njupt.edu.cn", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	context := string(result)
	if strings.Contains(context,"v46ip="){
		errReg, err := regexp.Compile("v46ip=([\\s\\S]*?);")
		if err != nil {
			fmt.Println(err)
		}
		return  strings.Split(errReg.FindString(context),"'")[1]
	}
	return ""
}

func DefaultLogout() {
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