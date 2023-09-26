#!/bin/sh

projectName=$1

mkdir $projectName

cd $projectName

go mod init $projectName

touch main.go