#!/bin/bash
rm libtun2socks.*
go build -o libtun2socks.a -buildmode=c-archive
