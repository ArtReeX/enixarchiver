package main

import (
	"fmt"
)

func getAnswer(variants []uint64) uint64 {

	for {

		// считывание строки
		var answer uint64
		fmt.Scanf("%d \n", &answer)

		// проверка допустимости ответа
		for _, variant := range variants {
			if variant == answer {
				return answer
			}
		}

		fmt.Println("Неправильный вариант ответа, попробуйте снова.")

	}

}

func main() {

	fmt.Println("Вы запустили архиватор для NieR: Automata.")
	fmt.Println("Выберите режим работы: 1 - распаковка | 2 - запаковка")

	getAnswer([]uint64{1, 2})

}
