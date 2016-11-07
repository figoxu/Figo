package Figo

import (
	"archive/zip"
	"io"
	"os"
)

func UnZip(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		defer rc.Close()
		if err != nil {
			return err
		}
		filename := dest + file.Name
		folderName := NewFilePath(filename).FolderName()
		if folderName == filename {
			continue
		}
		if err = os.MkdirAll(folderName, 0755); err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		if _, err = io.Copy(w, rc); err != nil {
			return err
		}
	}
	return nil
}
