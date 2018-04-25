package unpacking

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"

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
	buffer.Write(file)

	// чтение заголовков файла
	ReadHeader(&buffer)

}

// ReadHeader - функция для чтения заголовка файла
func ReadHeader(buffer *bytes.Buffer) *Header {

	// выделение места в памяти под заголовки
	head := Header{}

	fmt.Println("В памяти под заголовки выделено байт:", unsafe.Sizeof(head))

	// получение части, содержащую заголовок из буфера с автоматической прокруткой буфера
	partWithHead := buffer.Next(binary.Size(&head))

	fmt.Println("Заголовок для декодирования:", partWithHead)

	// преобразование данных заголовка в структуру
	if error := binary.Read(bytes.NewReader(partWithHead), binary.LittleEndian, &head); error != nil {
		fmt.Println("Ошибка преобразование данных заголовка в структуру.")
	}

	return &head

}
