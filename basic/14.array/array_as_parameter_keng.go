/*
初始化数组长度后，元素可以不进行初始化，或者不进行全部初始化，但未进行数组大小初始化的数组初始化结果元素大小就为多少。
*/
package main

import "fmt"

func main() {
    var array = []int{1, 2, 3, 4, 5}
    /* 未定义长度的数组只能传给不限制数组长度的函数 */
    setArray(array)
    /* 定义了长度的数组只能传给限制了相同数组长度的函数 */
    var array2 = [5]int{1, 2, 3, 4, 5}
    setArray2(array2)
}

func setArray(params []int) {
    fmt.Println("params array length of setArray is : ", len(params))
}

func setArray2(params [5]int) {
    fmt.Println("params array length of setArray2 is : ", len(params))
}
