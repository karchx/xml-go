package utils

import (
	"io/fs"
	"os"
)

func GetDateFile(file fs.FileInfo) string {
	return file.ModTime().Format(DefaultTimeFormat)
}

func GetNameFile(file os.File) string {
	return file.Name()
}
