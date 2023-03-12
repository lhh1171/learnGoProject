package myfile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Myfile6() {

	dat, err := ioutil.ReadFile("/MyProject/goProject/src/5project/xx/myfile.txt")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/MyProject/goProject/src/5project/xx/myfile.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//文件指针
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	//Read 从文件中读取最多 len(b) 个字节并将它们存储在 b 中。
	//它返回读取的字节数和遇到的任何错误。在文件末尾，Read 返回 0,
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	//ReadAtLeast 从 r 读取到 buf 中，直到它至少读取了 min 个字节。
	//如果读取的字节较少，它会返回复制的字节数和错误。仅当未读取任何字节时才会出现EOF错误。
	//如果在读取少于最小字节后发生EOF ，则 ReadAtLeast 返回ErrUnexpectedEOF 。
	//如果 min 大于 buf 的长度，ReadAtLeast 返回ErrShortBuffer 。
	//返回时，当且仅当 err == nil 时，n >= min。如果 r 返回读取至少 min 字节的错误，则错误将被丢弃。
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//Seek 将文件上的下一个Read或Write 的偏移量设置为 offset
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	//往后面读取5个
	//Peek 返回接下来的 n 个字节而不推进阅读器
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}
