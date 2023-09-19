package myupload

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// MyUpCli sample usage
func MyUpCli() {
	//此处为自己的ip地址
	target_url := "http://localhost:9090/upload"
	filename := "image2.jpg"
	postFile(filename, target_url)
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	//CreateFormFile 是CreatePart的便利包装器。
	//它使用提供的字段名称和文件名创建一个新的表单数据标题
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open("/MyProject/goProject/src/base5_io/mysock/mynet/myupload/" + filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy 把fh拷贝到fileWriter
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	//FormDataContentType 返回具有此Writer边界的HTTP multipart/form-data 的 Content-Type。
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	fmt.Println(bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
