package console

import (
	"bufio"
	"fmt"
	"os"
)

// GetNumber - функция получения цифры, выбранной пользователем, из строки
func GetNumber(variants []uint64) uint64 {

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

// Wait - функция ожидания нажатия пользователем любой клавиши
func Wait() {

	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

}
