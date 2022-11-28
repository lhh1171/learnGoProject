package gen

import (
	"fmt"
)

// CanSpeak
// 7. 类型+方法的约束
type CanSpeak interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
	Speak() string
}

type Mouth int32

func (m Mouth) Speak() string {
	return fmt.Sprintf("speak %v", m)
}

type Nose string

func (n Nose) Speak() string {
	return fmt.Sprintf("speak %v", n)
}

type Ear int32

func SpeakLoudly[T CanSpeak](params []T) {
	for _, param := range params {
		fmt.Println(param.Speak())
	}
}

func Test7() {
	// 7.1 类型与方法均符合
	SpeakLoudly([]Mouth{1, 2, 3, 4, 5, 6})

	// 7.2 类型符合
	//Nose的类型没有符合
	//./generics_test.go:172:16: Nose does not implement CanSpeak
	//SpeakLoudly([]Nose{"z", "x", "c"})

	// 7.3 方法符合
	//ear方法没有实现Speak
	//./generics_test.go:175:16: Ear does not implement CanSpeak (missing Speak method)
	//SpeakLoudly([]Ear{1, 2, 3, 4, 5, 6})
}
