package myssl

import (
	"io"
	"log"
	"net/http"
)

//openssl genrsa -out key.pem 2048
//openssl req -new -x509 -key key.pem -out cert.pem -days 3650

func MySsl() {
	http.HandleFunc("/hello", helloHandler)
	//cert.pem 客户端公钥
	//key.pem 服务端私钥
	err := http.ListenAndServeTLS(":8080",
		"/MyProject/goProject/src/goIO/mysock/mynet/myssl/cert.pem",
		"/MyProject/goProject/src/goIO/mysock/mynet/myssl/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
