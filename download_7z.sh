#!/bin/bash
set -ex

LIB7Z_VERSION="$(cat 7z.version)"
URL_7Z_WINDOWS_ALL=https://www.7-zip.org/a/7z${LIB7Z_VERSION}-extra.7z
URL_7Z_MAC_ALL=https://www.7-zip.org/a/7z${LIB7Z_VERSION}-mac.tar.xz
URL_7Z_LINUX_X86=https://www.7-zip.org/a/7z${LIB7Z_VERSION}-linux-x86.tar.xz
URL_7Z_LINUX_X86_64=https://www.7-zip.org/a/7z${LIB7Z_VERSION}-linux-x64.tar.xz
URL_7Z_LINUX_ARM=https://www.7-zip.org/a/7z${LIB7Z_VERSION}-linux-arm.tar.xz
URL_7Z_LINUX_ARM64=https://www.7-zip.org/a/7z${LIB7Z_VERSION}-linux-arm64.tar.xz

mkdir -p output /tmp/7z

echo "Downloading Windows binaries..."
wget "${URL_7Z_WINDOWS_ALL}" -O /tmp/7z/win32.7z
7z x /tmp/7z/win32.7z -o/tmp/7z/win32 -y
mv /tmp/7z/win32/x64/7za.exe output/7z_win32_x64.exe
mv /tmp/7z/win32/arm64/7za.exe output/7z_win32_arm64.exe
mv /tmp/7z/win32/License.txt output/7z-extra-license.txt
echo "Done."

echo "Downloading Linux and Mac binaries..."
mkdir -p /tmp/7z/linux-x86 /tmp/7z/linux-x64 /tmp/7z/linux-arm /tmp/7z/linux-arm64 /tmp/7z/mac
wget "${URL_7Z_LINUX_X86}" -O /tmp/7z/linux-x86.tar.xz
wget "${URL_7Z_LINUX_X86_64}" -O /tmp/7z/linux-x64.tar.xz
wget "${URL_7Z_LINUX_ARM}" -O /tmp/7z/linux-arm.tar.xz
wget "${URL_7Z_LINUX_ARM64}" -O /tmp/7z/linux-arm64.tar.xz
wget "${URL_7Z_MAC_ALL}" -O /tmp/7z/mac.tar.xz
tar xJf /tmp/7z/linux-x86.tar.xz -C /tmp/7z/linux-x86
tar xJf /tmp/7z/linux-x64.tar.xz -C /tmp/7z/linux-x64
tar xJf /tmp/7z/linux-arm.tar.xz -C /tmp/7z/linux-arm
tar xJf /tmp/7z/linux-arm64.tar.xz -C /tmp/7z/linux-arm64
tar xJf /tmp/7z/mac.tar.xz -C /tmp/7z/mac
mv /tmp/7z/mac/License.txt output/7z-unix-license.txt
mv /tmp/7z/linux-x86/7zz output/7z_linux_386
mv /tmp/7z/linux-x64/7zz output/7z_linux_x64
mv /tmp/7z/linux-arm/7zz output/7z_linux_arm
mv /tmp/7z/linux-arm64/7zz output/7z_linux_arm64
mv /tmp/7z/mac/7zz output/7z_darwin_arm64
cp output/7z_darwin_arm64 output/7z_darwin_x64
echo "Done."

echo "Cleanup..."
rm -rf /tmp/7z
echo "Done."