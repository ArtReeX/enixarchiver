package unpacking

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"../console"
	"../utilities"
)

// StartUnpackingMode - функция для запуска режима распаковки
func StartUnpackingMode(packedFolder string, unpackedFolder string) {

	fmt.Println(`
		Выбран режим распаковки. 
		Прямо сейчас положите файлы, которые необходимо расшифровать с расширением .dat в папку ` + packedFolder + `
		Затем нажмите любую клавишу.
	`)

	console.Wait()

	// получение всех файлов с папки
	listFiles := utilities.GetListFiles(packedFolder, []string{".dat"}, false)
	fmt.Println("Найдены следующие файлы для распаковки:", listFiles)

	// запуск процесса распаковки для каждого файла
	for _, file := range listFiles {
		fmt.Println("Начата распаковка файла:", file)
		Unpack(file)
	}

	fmt.Println(`
		Распаковка успешно завершена.
		Вы можете увидеть распакованные файлы в папке ` + unpackedFolder)

	console.Wait()

}

// Unpack - функция для распаковки файла
func Unpack(path string) {

	// чтение файла
	file := utilities.ReadFile(path)

	// создание потока для чтения файла
	buffer := bytes.Buffer{}
	binary.Write(&buffer, binary.BigEndian, file)

	// чтение заголовков файла
	ReadHeader(&buffer)

}

// ReadHeader - функция для чтения заголовка файла
func ReadHeader(buffer *bytes.Buffer) *Header {

	// выделение места в памяти под заголовки
	head := Header{}

	fmt.Println("В памяти под заголовки выделено байт:", binary.Size(&head))
	fmt.Println("Структура заголовка до декодирования:", head)

	// считывание заголовков
	WriteToStructUint32(buffer, &head.magic)
	WriteToStructUint32(buffer, &head.fileCount)
	WriteToStructUint32(buffer, &head.fileTableOffset)
	WriteToStructUint32(buffer, &head.extensionOffset)
	WriteToStructUint32(buffer, &head.nameTableOffset)
	WriteToStructUint32(buffer, &head.sizeTableOffset)
	WriteToStructUint32(buffer, &head.unknownOffset)
	WriteToStructUint32(buffer, &head.null)

	fmt.Println("Структура заголовка после декодирования:", head)

	return &head

}
