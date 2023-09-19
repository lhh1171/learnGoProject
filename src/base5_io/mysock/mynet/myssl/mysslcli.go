package myssl

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MySslCli() {
	//设置tls
	//InsecureSkipVerify 控制客户端是否验证服务器的证书链和主机名。
	//如果 InsecureSkipVerify 为真， 则 crypto/tls 接受服务器提供的任何证书以及该证书中的任何主机名。
	//在这种模式下，除非使用自定义验证， 否则 TLS 容易受到中间机器攻击。
	//这应该仅用于测试或与VerifyConnection或VerifyPeerCertificate结合使用。
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8080/hello")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
