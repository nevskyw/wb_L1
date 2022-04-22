package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.
*/

// https://go.dev/blog/slices-intro

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = v[:100]
}

func main() {
	v := createHugeString(1 << 10)      // 1024
	justString = strings.Clone(v[:100]) // выделит память под 100 элементов, а про 1024 он забудет
	fmt.Println(justString, len(justString))
}

func createHugeString(i int) string {
	var str string
	for i > 0 {
		str += "a"
		i--
	}
	return str
}

// func Clone(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	b := make([]byte, len(s))
// 	copy(b, s)
// 	return *(*string)(unsafe.Pointer(&b))
// }

//еще может быть речь про слайс по байтам а не по рунам
//да и в целом глобальные переменные для слабаков
/*
возможно речь идет про особенность GC в ГО — слайсы это всего лишь
сслыка (reference) на массив с которого взят слайс; а умный GC чистит данные
в том случае если к участку памяти больше никто не имеет доступа
Тем самым в данной реализации после создания большой строки, генерируется маленькая
которая является ссылкой на большой участок памяти. соответственно использование
маленькой строки влечет за собой СОХРАНЕНИЕ в памяти БОЛЬШОЙ потому что GC ее не
чистит. К тому же переменная глобальная и будет таскаться по всей программе.
*/

// Clone returns a fresh copy of s.
// It guarantees to make a copy of s into a new allocation,
// which can be important when retaining only a small substring
// of a much larger string. Using Clone can help such programs
// use less memory. Of course, since using Clone makes a copy,
// overuse of Clone can make programs use more memory.
// Clone should typically be used only rarely, and only when
// profiling indicates that it is needed.
// For strings of length zero the string "" will be returned
// and no allocation is made.
