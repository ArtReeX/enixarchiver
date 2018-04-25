package unpacking

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// WriteToStructUint32 - функция для записи данных с буфера в элемент структуры
func WriteToStructUint32(buffer *bytes.Buffer, element *uint32) {

	part := buffer.Next(binary.Size(element))

	if error := binary.Read(bytes.NewReader(part), binary.BigEndian, element); error != nil {
		fmt.Println("Ошибка записи с буфера в структуру.", error)
	}

}
