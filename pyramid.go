package main

import (
  "fmt"
  "bufio"
  "os"
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
}

type pyramid struct {
  //*root, *temp , *node_buff node struct
  gameover bool
  root,temp,node_buff  node
  win[6299] uint
}

func (py pyramid ) display(){
  p := fmt.Println
  p("   *   ")
  p("  ***  ")
  p(" ***** ")
  p("*******")
}

func (py pyramid ) you_move(){
  p := fmt.Println
  p("you move")
  py.display()
}

func (py pyramid) my_move(){
  p := fmt.Println
  p("my move")
  py.display()
}

func (py pyramid) is_gameover() bool {
  return py.gameover
}

func main(){
  py := pyramid{gameover: false}

	fmt.Println("Welcome to Pyramid game ")
	fmt.Println("The Rules of Pyramid: Players alternate take out one to three ")
	fmt.Println("tokens in one line. Who takes the last one token is loser")

  for {
    py.gameover = false
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Do you want to move first(y/n)?")
    ans,_ := reader.ReadByte()
    _,_ = reader.ReadByte()
    //strings.Replace(ans,"\r\n","",-1)

    fmt.Println(ans)
    fmt.Println("done")

    if ans == 'y'  {
      fmt.Println(" OK ! You move first.")
      for !py.is_gameover() {
        py.you_move();
        py.gameover = true
        if py.is_gameover() {
          fmt.Println("I win")
        } else {
          py.you_move()
          if py.is_gameover() {
            fmt.Println("You win")
          }
        }
      }
    } else {
      fmt.Println(" OK ! I  move first.")
      for !py.is_gameover() {
        py.my_move()
        py.gameover = true
        if py.is_gameover() {
          fmt.Println("You win")
        } else  {
          py.you_move()
          if py.is_gameover() {
            fmt.Println("I win")
          }
        }
      }
    }

    fmt.Println("Game Over")
    fmt.Print("Do you want to play again(y/n)?")
    ans,_ = reader.ReadByte()
    _,_ = reader.ReadByte()
   // strings.Replace(ans,"\r","",-1)
    fmt.Println(ans)

    if  ans == 'n'   { break }
  }

}
