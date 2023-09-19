package myapi

import "crypto/sha1"
import "fmt"

func MySha1() {
	s := "sha1 this string"

	/*包 sha1 实现了 RFC 3174 中定义的 SHA-1 散列算法。
	SHA-1 是密码破损的，不应用于安全应用程序*/
	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
