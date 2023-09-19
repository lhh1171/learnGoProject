package mynet

import (
	"io"
	"log"
	"net/http"
)

func MyHttp22() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler22{})
	mux.HandleFunc("/hello", sayhello22)

	wd := "/MyProject/goProject/src/base5_io/sss/"

	mux.Handle("/sss/", http.StripPrefix("/sss/", http.FileServer(http.Dir(wd))))

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler22 struct {
}

func (this *myHandler22) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world!")
}

func sayhello22(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world! 21")
}
