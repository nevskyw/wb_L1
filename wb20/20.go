package main

import "fmt"

func main() {
	newMessageString := "snow dog sun"

	mapa := make(map[int]string)

	num := 0 // нужен для ключа
	var word string

	// Пробегаем по переменной newMessageString обращаясь к каждому символу отдельно, пока мы не береберем все полностью
	for _, i := range newMessageString { 

		if string(i) != " " { // Если i не равен пробелу в строке
			word += string(i) // Складываем буквы с новой переменной
			
		} else {
			mapa[num] = word // Записываем переменную word в карту messageSlice с ключом num
			word = ""                // Очищаем переменную 
			num += 1                 // Увеличиваем счетчик, чтобы при записи в карту - ключи отличались
		}
	}
	mapa[num] = word // Записываем последнее слово в карту

	var endMessage string // создаем переменную для выводы в консоли

	// Перебираем все значения в карте и прибавляем их к новой переменной, так же не забываем прибавить пробел
	for i := 0; i < len(mapa); i++ {
		endMessage = mapa[i] + " " + endMessage // слова прибавляются справо на лево
	}

	fmt.Println(endMessage) // Выводим перевернутую строку
}
