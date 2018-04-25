package main

import (
	"fmt"

	"./console"
	"./unpacking"
)

const (
	// PACKEDFOLDER - папка с запакованными файлами
	PACKEDFOLDER = "packed"
	// UNPACKEDFOLDER - папка с распакованными файлами
	UNPACKEDFOLDER = "unpacked"
)

// StartMode - функция запуска выбранного пользователем режима
func StartMode(mode uint64) {

	switch mode {

	case 1:
		unpacking.StartUnpackingMode(PACKEDFOLDER, UNPACKEDFOLDER)
	}

}

func main() {

	fmt.Println(`
		Вы запустили архиватор для NieR: Automata.
		Выберите режим работы: 1 - распаковка | 2 - запаковка.
	`)

	StartMode(console.GetNumber([]uint64{1, 2}))

}
