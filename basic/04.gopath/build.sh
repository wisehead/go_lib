#############################################################
#   File Name: build.sh
#     Autohor: Hui Chen (c) 2022
#        Mail: alex.chenhui@gmail.com
# Create Time: 2022/02/14-09:32:57
#############################################################
#!/bin/sh 
go env -w GO111MODULE=off
cp -R ./src/proj/myMath $GOPATH/src/proj
cd ./src/proj/test
go build test.go
go env -w GO111MODULE=on
