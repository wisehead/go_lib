#############################################################
#   File Name: build.sh
#     Autohor: Hui Chen (c) 2020
#        Mail: chenhui13@baidu.com
# Create Time: 2020/04/09-21:38:56
#############################################################
#!/bin/sh 
gcc -c hello.c
ar -cru libhello.a hello.o
