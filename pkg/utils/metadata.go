package utils

import (
	"io/fs"
	"os"
)

func GetDateFile(file fs.FileInfo) string {
	return file.ModTime().Format("DD/MM/YYYY hh:mm")
}

func GetNameFile(file os.File) string {
	return file.Name()
}
