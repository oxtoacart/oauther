#!/bin/bash

# This script builds binaries for windows, linux and mac, compresses them with
# upx and puts them in the dist folder.
source ~/git/golang-crosscompile/crosscompile.bash
mkdir dist
go-windows-386 build -o dist/oauther_i386.exe
go-windows-amd64 build -o dist/oauther_amd64.exe
go-linux-386 build -o dist/oauther_linux_i386
go build -o dist/oauther_osx

upx dist/*
