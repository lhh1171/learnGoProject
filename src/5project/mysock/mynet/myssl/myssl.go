package myssl

import (
	"io"
	"log"
	"net/http"
)

//openssl genrsa -out key.pem 2048
//openssl req -new -x509 -key key.pem -out cert.pem -days 3650

// MySsl
//1、生成RSA密钥的方法
//openssl genrsa -des3 -out key.pem 2048
//这个命令会生成一个2048位的密钥，同时有一个des3方法加密的密码，如果你不想要每次都输入密码，可以改成：
//openssl genrsa -out key.pem 2048
//建议用2048位密钥，少于此可能会不安全或很快将不安全。
//
//2、生成一个证书请求
//openssl req -new -key key.pem -out cert.csr
//这个命令将会生成一个证书请求，当然，用到了前面生成的密钥key.pem文件
//这里将生成一个新的文件cert.csr，即一个证书请求文件，你可以拿着这个文件去数字证书颁发机构（即CA）申请一个数字证书。CA会给你一个新的文件cert.pem，那才是你的数字证书。
//
//如果是自己做测试，那么证书的申请机构和颁发机构都是自己。就可以用下面这个命令来生成证书：
//openssl req -new -x509 -key key.pem -out cert.pem -days 1095
//这个命令将用上面生成的密钥key.pem生成一个数字证书cert.pem
//
//3、使用数字证书和密钥
//有了key.pem和cert.pem文件后就可以在自己的程序中使用了，比如做一个加密通讯的服务器/*
func MySsl() {
	http.HandleFunc("/hello", helloHandler)
	//cert.pem 客户端公钥
	//key.pem 服务端私钥
	//ListenAndServeTLS 的行为与ListenAndServe相同，只是它需要 HTTPS 连接。
	//此外，必须提供包含服务器证书和匹配私钥的文件。
	//如果证书由证书颁发机构签署，则 certFile 应该是服务器证书、任何中间证书和 CA 证书的串联
	err := http.ListenAndServeTLS(":8080",
		"/MyProject/goProject/src/5project/mysock/mynet/myssl/cert.pem",
		"/MyProject/goProject/src/5project/mysock/mynet/myssl/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
