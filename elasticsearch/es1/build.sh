#############################################################
#   File Name: build.sh
#     Autohor: Hui Chen (c) 2022
#        Mail: alex.chenhui@gmail.com
# Create Time: 2022/02/17-19:09:06
#############################################################
#!/bin/sh 
go mod init es1
go mod tidy
go build es.go
curl 'localhost:9200/user/_search'
