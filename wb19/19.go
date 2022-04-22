package main

import "fmt"

func Reverse(s string) (result string) { // объявляем переменную result в возыращаемом значении
	for _, v := range s {
		result = result + string(v) // буквы прибавляются справо на лево
		fmt.Println(result)
	}
	return
}
func main() {
	fmt.Println(Reverse("Golang"))
}
