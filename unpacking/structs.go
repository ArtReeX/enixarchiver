package unpacking

// Header - структура файла
type Header struct {
	magic           uint32 // 'DAT\x00'
	fileCount       uint32
	fileTableOffset uint32
	extensionOffset uint32
	nameTableOffset uint32
	sizeTableOffset uint32
	unknownOffset   uint32
	null            uint32 // Zero
}

// FileEntry - структура файла
type FileEntry struct {
	offset uint32
}

// ExtentionTableEntry - структура таблицы
type ExtentionTableEntry struct {
	extension [4]rune // 'bin\x00' , 'pak\x00', 'z\x00\x00\x00', etc
	u32       uint32
}

// FileSizeEntry - структура размера
type FileSizeEntry struct {
	size uint32
}
