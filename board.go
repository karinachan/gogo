package main

import "fmt"

type Board [10][10]int

func (b Board) toString() string{
  str := "["
  for i, elem := range b {
          if i > 0 {
              str += " "
          }
          str += fmt.Sprint(elem)
          str += "\n"
      }
      return str + "]"

}

func main() {
    b := Board{}
    fmt.Println("board contents: \n", b.toString())
}
