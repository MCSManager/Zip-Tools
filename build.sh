#!/bin/bash
set -ex

mkdir -p output

export CGO_ENABLED=0

GOOS=windows GOARCH=amd64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_win32_x64.exe
GOOS=linux GOARCH=amd64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_x64
GOOS=linux GOARCH=arm64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_arm64
GOOS=linux GOARCH=arm go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_arm
GOOS=linux GOARCH=386 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_386
GOOS=linux GOARCH=mips go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_mips
GOOS=linux GOARCH=mips64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_mips64
GOOS=linux GOARCH=mipsle go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_mipsle
GOOS=linux GOARCH=ppc64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_ppc64
GOOS=linux GOARCH=riscv64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_riscv64
GOOS=linux GOARCH=s390x go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_linux_s390x
GOOS=netbsd GOARCH=amd64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_netbsd_x64
GOOS=netbsd GOARCH=arm go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_netbsd_arm
GOOS=netbsd GOARCH=arm64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_netbsd_arm64
GOOS=openbsd GOARCH=386 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_openbsd_386
GOOS=openbsd GOARCH=amd64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_openbsd_x64
GOOS=openbsd GOARCH=arm go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_openbsd_arm
GOOS=openbsd GOARCH=arm64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_openbsd_arm64
GOOS=freebsd GOARCH=386 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_freebsd_386
GOOS=freebsd GOARCH=amd64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_freebsd_x64
GOOS=freebsd GOARCH=arm go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_freebsd_arm
GOOS=freebsd GOARCH=arm64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_freebsd_arm64
GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_darwin_amd64
GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags '-s -w --extldflags "-static -fpic"' -o output/file_zip_darwin_arm64