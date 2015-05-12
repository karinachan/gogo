package main

import "fmt"


type Board[9][9]int
type Turn[3]int

func (b Board) toString() string{
  str := "[  0 1 2 3 4 5 6 7 8 \n "
  for i, elem := range b {
          if i > 0 {
              str += " "
          }
          str += fmt.Sprint(i)
          str += " "
          str += fmt.Sprint(elem)
          str += "\n"
      }
      return str + "]"

}

func (b Board) setPiece(row int, col int, piece int) Board{
  b[row][col] = piece
  return b
}

func (b Board) getPiece(row int, col int) int {
  return b[row][col]
}

func (t Turn) setTurn(row int, col int, piece int) Turn {
  t[0] = row
  t[1] = col
  t[2] = piece
  return t
}


  // if piece is row is 0 or 9 or col is 0 or 9 check 3
  // if piece is row 0/9 col 0/9 check 2
  // else check 4
 /* var i = b[row][col]
  var x[4]int 
  if row == 0 { 
    if col == 0 { // 0,0
      x[0] = 0
      x[3] = 0
      if b[row][col+1] == i {

      }
      if b[row+1][col] == i {

      }
    } else if col == 8 { // 0,8
      x[0] = 0
      x[1] = 0 
      if b[row][col-1] == i {

      }
      if b[row+1][col] == i {

      }
    } else { // 0,y
        x[0] = 0
        if b[row][col+1] == i {

        }
        if b[row][col-1] == i {

        }
        if b[row+1][col] == i {

        }
    }   
  } else if row == 8 { 
    if col == 0 { // 8,0
      x[2] = 0
      x[3] = 0
      if b[row-1][col] == i {

      }
      if b[row][col+1] == i {

      }
    } else if col == 8 { // 8,8
      x[1] = 0
      x[2] = 0
      if b[row-1][col] == i {

      }
      if b[row][col-1] == i {
        
      }
    } else { // 8,y
      x[2] =0
      if b[row][col+1] == i {

        }
        if b[row][col-1] == i {

        }
        if b[row-1][col] == i {

        }
    }
  } else {
    if col == 0 { // x,0
      x[3] = 0
      if b[row][col+1] == i {

        }
        if b[row+1][col] == i {

        }
        if b[row-1][col] == i {

        }
    } else if col == 8 { // x,8
      x[1] = 0
        if b[row][col-1] == i {

        }
        if b[row+1][col] == i {

        }
        if b[row-1][col] == i {

        }
    } else { // x,y
         if b[row+1][col] == i {

        }
        if b[row][col+1] == i {

        }
        if b[row-1][col] == i {

        } 
        if b[row][col-1] == i {

        }
    }   
  } */
  //FLOODFILL
  func (b Board) isLoop(row int, col int, t Turn, ) bool {
  var i = b.getPiece(row,col)
  if row == t[0] && col == t[1] && i == t[2] {
    //add to list of loops
    //take longest of loops
    //loop can be a single piece, or 2 length path
    // break this recursion
  } 

  // iterate through every direction
  // how to break infinite recursion?
  
  if b.getPiece(row+1,col) == i {
    b.isLoop(row+1,col,t)
  }
  
  if b.getPiece(row,col+1) == i {
 
    Loop(row,col+1,t)
  }
  
  if b.getPiece(row-1,col) == i {
    b.isLoop(row-1,col,t)
  }
  
  if b.getPiece(row,col-1) == i {
    b.isLoop(row,col-1,t)
  }

  return true
}



func main() {
    b := Board{}
    t := Turn{}
    b = b.setPiece(2,2,1)
    t = t.setTurn(2,2,1)
    fmt.Println(fmt.Sprint(t))
    b.isLoop(2,2,t)
    fmt.Println("board contents: \n", b.toString())
}
