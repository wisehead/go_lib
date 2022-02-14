package main

const (
    i = iota
    j = iota
    x = iota
)
const xx = iota
const yy = iota
func main(){
    println(i, j, x, xx, yy)
}

// 输出是 0 1 2 0 0
