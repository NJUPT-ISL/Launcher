package operations

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func PostOperation(user string, pass string) {
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
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
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
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}