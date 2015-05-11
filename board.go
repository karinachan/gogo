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

func (b Board) setPiece(row int, col int, piece int) Board{
  b[row][col] = piece; 
   //   fmt.Println("board contents: \n", b.toString())
  return b

}

func main() {
    b := Board{}
    b = b.setPiece(1,1,1)
    fmt.Println("board contents: \n", b.toString())

}
