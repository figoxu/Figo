package Figo

import (
	zipTool "github.com/pierrre/archivefile/zip"
	"log"
	"os"
)

func Zip(dir, dst string) error {
	outFile, err := NewFilePath(dst).Open()
	if err != nil {
		return err
	}
	defer outFile.Close()
	progress := func(archivePath string) {
		log.Println(archivePath)
	}
	return zipTool.Archive(dir, outFile, progress)
}

func UnZip(zipFile, destDir string) error {
	if err := os.MkdirAll(destDir, 0777); err != nil {
		return err
	}
	progress := func(archivePath string) {
		log.Println(archivePath)
	}
	return zipTool.UnarchiveFile(zipFile, destDir, progress)
}
