#############################################################
#   File Name: build.sh
#     Autohor: Hui Chen (c) 2022
#        Mail: alex.chenhui@gmail.com
# Create Time: 2022/02/11-16:19:13
#############################################################
#!/bin/sh 
go mod init xxx
cd test
go build test.go
./test
