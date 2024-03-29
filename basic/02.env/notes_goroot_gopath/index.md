
# go环境变量配置 (GOROOT和GOPATH) - 简书

* * *

1.  GOROOT就是go的安装路径  
    在~/.bash\_profile中添加下面语句:

```bash
GOROOT=/usr/local/go
export GOROOT
```

当然, 要执行go命令和go工具, 就要配置go的可执行文件的路径:  
操作如下:  
在~/.bash\_profile中配置如下:  
`export $PATH:$GOROOT/bin`  
如果是windows需要使用`;`符号分割两个路径, mac和类unix都是用`:`符号分割

2.  GOPATH:

*   `go install`/`go get`和 go的工具等会用到GOPATH环境变量.
*   GOPATH是作为编译后二进制的存放目的地和import包时的搜索路径 (其实也是你的工作目录, 你可以在src下创建你自己的go源文件, 然后开始工作)。
    1.  GOPATH之下主要包含三个目录: bin、pkg、src
    2.  bin目录主要存放可执行文件; pkg目录存放编译好的库文件, 主要是\*.a文件; src目录下主要存放go的源文件
*   不要把GOPATH设置成go的安装路径,  
    可以自己在用户目录下面创建一个目录, 如gopath  
    操作如下:

```bash
cd ~
mkdir gopath
```

在~/.bash\_profile中添加如下语句:  
`GOPATH=/Users/username/gopath`

*   GOPATH可以是一个目录列表, `go get`下载的第三方库, 一般都会下载到列表的第一个目录里面
*   需要把GOPATH中的可执行目录也配置到环境变量中, 否则你自行下载的第三方go工具就无法使用了, 操作如下:  
    在~/bash\_profile中配置,

```bash
export $PATH:$GOPATH/bin
```

*   创建一个go项目, 并且编译运行:

```bash
mkdir goproject
cd goproject
touch hello.go
```

在hello.go中输入:

```go
package main
import "fmt"
func main() {
      fmt.Println("Hello, GO !")
}
```

在项目根目录下执行`go build`命令来构建你的项目, 构建后会生成hello文件  
运行生成的文件`./hello`, terminal中输出: `Hello, GO !`  
当然你也可以直接运行命令`go run hello.go`来执行程序.

_如果要上述设置生效, 可以执行命令: `source ~/.bash_profile`, 上述所有操作均为mac系统下的操作, 如果是非mac系统, 请自行变通._