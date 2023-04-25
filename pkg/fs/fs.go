package fs

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/gothew/l-og"
	"github.com/karchx/xml-go/pkg/utils"
)

func MoveMatchNewFolder(src string) {
	originalFileName := strings.Split(src, "/")

	if _, err := os.Stat(utils.DefaultDestionDirectory); os.IsNotExist(err) {
		err := os.Mkdir(utils.DefaultDestionDirectory, 0777)
		if err != nil {
			log.Fatalf("Error create directory: %s", err)
		}
	}

	err := copyMatch(src, utils.DefaultDestionDirectory+"/"+originalFileName[1])
	if err != nil {
		log.Fatalf("Error move file: %s", err)
	}
}

func copyMatch(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	log.Infof("Copy from %s to %s", src, dst)
	return err
}
