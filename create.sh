#!/bin/bash

if [ ! -f go.mod ];then go mod init chat-room;fi
if [ ! -d deploy ];then mkdir deploy;fi

if [ ! -d deploy/config ];then mkdir deploy/config;fi
if [ ! -d deploy/logs ];then mkdir deploy/logs;fi
if [ ! -d deploy/doc ];then mkdir deploy/doc;fi
if [ ! -d deploy/static ];then mkdir deploy/static;fi
if [ ! -d deploy/static/file ];then mkdir deploy/static/file;fi

base_dir=$(pwd)

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
