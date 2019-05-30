package core

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"xpath"
)

func createFile(filename, fileContext string) {
	filePoint, err := os.Create(filename)
	fmt.Println(filename)
	if err != nil {
		log.Println("os.Create(filename)")
		log.Fatal(err)
	}
	defer filePoint.Close()
	Writer := bufio.NewWriter(filePoint)
	Writer.WriteString(fileContext)
	Writer.Flush()
}

func (basic *Basic) Checker() (bool, error) {
	if basic.RemotePath != "" && basic.IndexFileName != "" && basic.LocalPath != "" {
		remoteURL, localURL := basic.RemotePath+basic.IndexFileName, basic.LocalPath+basic.IndexFileName
		remoteRelease, localRelease := xpath.Request(remoteURL), xpath.Request(localURL)
		remoteFileName, localFileName := "remote_"+basic.ImageName, "local_"+basic.ImageName
		createFile(remoteFileName, remoteRelease)
		createFile(localFileName, localRelease)
		if sumFileMd5(remoteFileName) == sumFileMd5(localFileName) {
			return true, nil
		}
	} else {
		err := errors.New("RemotePath/ is nil")
		return false, err
	}
	return false, nil
}

func sumFileMd5(fileName string) string {
	filePoint, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer filePoint.Close()
	reader := bufio.NewReader(filePoint)
	hasher := md5.New()
	_, err = io.Copy(hasher, reader)
	if err != nil {
		log.Println("io.Copy(hasher, reader)")
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
