package main

import (
  "fmt"
  "bufio"
  "os"
  "math"
  //"strings"
)


type node struct {
  step      uint
  bestchild uint
  children  uint
  cnum      int
  value     int
}

type game interface {
  display()
  you_move()
  my_move()
  is_gameover() bool
  is_valid_move(string) bool
}


type board struct {
  vals [15] int
}

type pyramid struct {
  //*root, *temp , *node_buff node struct
  //board    [15] int
  board board
  root,temp,node_buff  node
  win[6299] uint
}

func (b board) get(pos int) int {
  return b.vals[pos]
}

func (b *board) set(pos int) {
  b.vals[pos] = 1
}

func (b board) val() int {
  var sum int
  for i,r := range b.vals {
    sum +=  r * int(math.Pow(2,float64(i)) )
  }
  return sum
}


//  0(0,4)
//  1(1,3) 2(1,5) 
//  3(2,2) 4(2,4) 5(2,6)
//  6(3,1) 7(3,3) 8(3,5) 9(3,7)
//  A(4,0) B(4,2) C(4,4) D(4,6) E(4,8)
func (py pyramid ) display(){
  // alias
  pl := fmt.Println
  p := fmt.Print
  d := 5  // depth 
  t :=  0 // start token value 
  for i:= 0 ; i < d  ; i++ {
    // print prefix space
    for k := 1 ; k < d - i ; k++ {
      p(" ")
    }

    // print token
    for j :=0 ; j <  i+1 ; j ++ {
      if py.board.get(t) == 0 {
        p(string(t+65))
      } else {
        p("*")
      }
      p(" ")
      t++;
    }

    // print endl
    pl("")
  }
  pl( py.board.val())
}

func (py *pyramid ) you_move(){
  p  := fmt.Print
  p("Please enter your move (A-N)? ")
  reader := bufio.NewReader(os.Stdin)
  var str string
  for {
    str,_ = reader.ReadString('\n')
    str = str[:len(str)-1]

    // only allow pick 1,2,3 token
    if len(str) > 0 && len(str) <= 3 { 
      break
    }

    // move validation
    if py.is_valid_move(str) {
      break
    }

    p("Invalid !!! Please enter your move again (A-N)? ")
  }

  for _, r := range str {
    py.board.set( int(r - 65) )
  }

  py.display()
}

func (py pyramid) is_valid_move ( str string) bool {
  return true
}

func (py *pyramid) my_move(){
  p := fmt.Println
  p("my move")
  py.display()
}

func (py pyramid) is_gameover() bool {
  for i:=0 ; i < 15 ; i++ {
    if py.board.get(i) == 0 {
      return false
    }
  }
  return true
}

func main(){
  pl := fmt.Println

  py := pyramid{}

	pl("Welcome to Pyramid game ")
	pl("The Rules of Pyramid: Players alternate take out one to three ")
	pl("tokens in one line. Who takes the last one token is loser")

  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Do you want to move first(y/n)?")
    ans,_ := reader.ReadByte()
    _,_ = reader.ReadByte()

    if ans == 'y'  {
      py.display()
      pl("OK ! You move first.")
      for !py.is_gameover() {
        py.you_move();
        if py.is_gameover() {
          pl("I win")
        } else {
          py.you_move()
          if py.is_gameover() {
            pl("You win")
          }
        }
      }
    } else {
      pl(" OK ! I  move first.")
      for !py.is_gameover() {
        py.my_move()
        if py.is_gameover() {
          pl("You win")
        } else  {
          py.you_move()
          if py.is_gameover() {
            pl("I win")
          }
        }
      }
    }

    pl("Game Over")
    pl("Do you want to play again(y/n)?")
    ans,_ = reader.ReadByte()
    _,_ = reader.ReadByte()

    if  ans == 'n'   { break }
  }

}
