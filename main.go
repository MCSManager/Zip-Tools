package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/transform"
)

var zipPath string
var distDirPath string
var mode int
var encode string

const (
	ZIP_MODE   int = 1
	UNZIP_MODE int = 2
)

type SrcFiles []string

func (i *SrcFiles) String() string {
	return ""
}

func (i *SrcFiles) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var srcFiles SrcFiles

func init() {
	flag.StringVar(&encode, "code", "", "support GBK,UTF-8")
	flag.StringVar(&zipPath, "zipPath", "", "zip file path")
	flag.StringVar(&distDirPath, "distDirPath", "", "dir path")
	flag.Var(&srcFiles, "file", "-file 1.txt -file 2.txt -file 3.txt")
	flag.IntVar(&mode, "mode", ZIP_MODE, "1=zip, 2=unzip")
}

func main() {
	flag.Parse()
	if mode == UNZIP_MODE {
		fmt.Println("UNZIP:", zipPath, "-->", distDirPath, "code:", encode)
		err := UnZip(distDirPath, zipPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(-3)
		}
		return
	}
	if mode == ZIP_MODE {
		fmt.Println("ZIP:", srcFiles, "-->", zipPath, "code:", encode)
		err := Zip(srcFiles, zipPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(-2)
		}
		return
	}
	fmt.Println("Err Mode Params!")
	os.Exit(-1)
}

func Zip(srcFiles []string, dest string) error {
	var files []*os.File
	for _, fileName := range srcFiles {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		files = append(files, file)
	}
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := fileToZipWriter(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnZip(distDirPath, zipPath string) error {
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}

	defer zipFile.Close()

	prefix := ""
	if distDirPath != "" && distDirPath != "." {
		prefix = distDirPath + "/"
	}

	for _, f := range zipFile.File {
		nameReader := bytes.NewReader([]byte(f.Name))
		decoder := transform.NewReader(nameReader, getDecoderByCoder(encode))
		content, _ := ioutil.ReadAll(decoder)
		filePath := string(content)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(prefix+filePath, os.ModePerm)
			continue
		}
		if err := os.MkdirAll(prefix+filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}
		dstFile, err := os.OpenFile(prefix+filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		file, err := f.Open()
		if err != nil {
			return err
		}
		if _, err := io.Copy(dstFile, file); err != nil {
			return err
		}
		dstFile.Close()
		file.Close()
	}
	return nil
}
