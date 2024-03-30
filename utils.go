package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

func getDecoderByCoder(code string) *encoding.Decoder {
	code = strings.ToLower(code)
	var decoder *encoding.Decoder
	switch code {
	case "gbk":
		decoder = simplifiedchinese.GBK.NewDecoder()
	case "big5":
		decoder = traditionalchinese.Big5.NewDecoder()
	case "shift_jis":
		decoder = japanese.ShiftJIS.NewDecoder()
	case "euckr":
		decoder = korean.EUCKR.NewDecoder()
	default:
		decoder = unicode.UTF8.NewDecoder()
	}
	return decoder
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
