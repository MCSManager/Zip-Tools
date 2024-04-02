# File-Zip

A lightweight compression and decompression tool.

## Features

Unzip a ZIP file

```bash
./file-zip -mode 2 --zipPath MyZip.zip --DistDirPath MyDir --code utf-8
```

- `--code`: `utf-8`, `gbk`, `big5`, `shift_jis`, `euckr`

Compress multiple files

```bash
./file-zip -mode 1 --file test1.txt --file TestDir --file test2.txt --zipPath NewZip.zip
```

### License

Released under the MIT License.
