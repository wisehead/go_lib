package main

import "fmt"

func main() {
  arr := [...][]int{
    {1, 2, 3, 4},
    {10, 20, 30, 40},
  }
  for i := range arr {
    for j := range arr[i] {
        fmt.Println(arr[i][j])
    }
  }
}
