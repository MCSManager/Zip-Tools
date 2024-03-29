package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var zipPath string
var distDirPath string
var mode int

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
	flag.StringVar(&zipPath, "zipPath", "", "zip file path")
	flag.StringVar(&distDirPath, "DistDirPath", "", "dir path")
	flag.Var(&srcFiles, "file", "--file 1.txt --file 2.txt --file 3.txt")
	flag.IntVar(&mode, "mode", ZIP_MODE, "1=zip, 2=unzip")
}

func main() {
	flag.Parse() // 解析参数

	fmt.Println("PPPP:", zipPath, distDirPath, mode, srcFiles)

	if mode == UNZIP_MODE {
		fmt.Println("UNZIP:", zipPath, "-->", distDirPath)
		UnZip(distDirPath, zipPath)
		return
	}
	if mode == ZIP_MODE {
		fmt.Println("ZIP:", srcFiles, "-->"+zipPath)
		Zip(srcFiles, zipPath)
		return
	}

	fmt.Println("ERR Params!")
	os.Exit(-1)
}

// 压缩文件
// files 文件数组，可以是不同dir下的文件或者文件夹
// dest 压缩文件存放地址
func Zip(srcFiles []string, dest string) error {
	var files []*os.File
	for _, fileName := range srcFiles {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalln(err)
			os.Exit(-4)
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

func fileToZipWriter(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if prefix == "" {
			prefix = info.Name()
		} else {
			prefix = prefix + "/" + info.Name()
		}

		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = fileToZipWriter(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if prefix != "" {
			header.Name = prefix + "/" + header.Name
		}
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func UnZip(distDirPath, zipPath string) {
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	prefix := ""
	if distDirPath != "" && distDirPath != "." {
		prefix = distDirPath + "/"
	}

	// 第二步，遍历 zip 中的文件
	for _, f := range zipFile.File {
		filePath := f.Name
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(prefix+filePath, os.ModePerm)
			continue
		}
		// 创建对应文件夹
		if err := os.MkdirAll(prefix+filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}
		// 解压到的目标文件
		dstFile, err := os.OpenFile(prefix+filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}
		file, err := f.Open()
		if err != nil {
			panic(err)
		}
		// 写入到解压到的目标文件
		if _, err := io.Copy(dstFile, file); err != nil {
			panic(err)
		}
		dstFile.Close()
		file.Close()
	}
}
