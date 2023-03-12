package bytes

import (
	"bytes"
	"fmt"
)

func BytesBuffer() {

	b1 := []byte("hello world!")
	buf := bytes.NewBuffer(b1)
	fmt.Printf("buff len=%d\n", buf.Len())
	fmt.Printf("buff cap=%d\n", buf.Cap())

	buf.Grow(100)
	fmt.Printf("buff len=%d\n", buf.Len())
	fmt.Printf("buff cap=%d\n", buf.Cap())

	b2 := make([]byte, 6)
	// Read 从缓冲区中读取下一个 len(p) 个字节，或者直到缓冲区被耗尽。返回值 n 是读取的字节数。
	//如果缓冲区没有数据可返回，则 err 为 io.EOF（除非 len(p) 为零）；否则为零。
	buf.Read(b2)
	println(string(b2)) //hello

	b3 := buf.Next(5)
	println(string(b3)) //world

	b4 := buf.Next(3)
	println(string(b4))

	buf2 := bytes.NewBuffer(b1)
	// ReadBytes 一直读取到输入中第一次出现分隔符，返回一个包含数据的切片，直到并包括分隔符。
	b5, _ := buf2.ReadBytes(byte(' '))
	println(len(b5))
	println(string(b5))

	b6 := []byte("go programming")
	buf3 := bytes.NewBuffer(b1)
	// Write 将 p 的内容附加到缓冲区，根据需要增大缓冲区。
	buf3.Write(b6)
	println(string(buf3.Bytes()))
}
