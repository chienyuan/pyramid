package main

import (
  "fmt"
  "bufio"
  "os"
  "math"
  "sort"
  //"strings"
)


type node struct {
  step      uint
  bestchild uint
  children  uint
  cnum      int
  value     int
}

// token vals 
// 0 : free
// 1 : non-free
type board struct {
  vals [15] int
}

type pyramid struct {
  //*root, *temp , *node_buff node struct
  //board    [15] int
  board board
  root,temp,node_buff  node
  win[6299] uint
  valid_moves[] int
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
  pl := fmt.Println
  p  := fmt.Print
  py.gen_valid_moves()
  p("Please enter your move (A-N)? ")
  reader := bufio.NewReader(os.Stdin)
  var str string
  for {
    str,_ = reader.ReadString('\n')
    str = str[:len(str)-1]

    pl("len(str)=",len(str))
    // only allow pick 1,2,3 token
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
    if len(str) < 0 || len(str) > 3 { 
      return false
    }

    var move int
    for _, r := range str {
      move += int( math.Pow(2,float64(r - 65) ) )
    }
    fmt.Println("move=",move)

    for _,r := range py.valid_moves {
      if move == r {
        return true
      }
    }

  return false
}

func (py *pyramid) gen_valid_moves() {
  //    0
  //   1 2
  //  3 4 5
  // 6 7 8 9
  //a b c d e
  // two token move data
  var arr[63] int
	var temp2 = [30][2]  int {
	{0,1},{1,3},{3,6},{6,10},{2,4},{4,7},{7,11},{5,8},{8,12},{9,13},
	{0,2},{2,5},{5,9},{9,14},{1,4},{4,8},{8,13},{3,7},{7,12},{6,11},
	{1,2},{3,4},{4,5},{6,7},{7,8},{8,9},{10,11},{11,12},{12,13},{13,14}}

	// three token move data
	var temp3 = [18][3] int {
		{0,1,3},{1,3,6},{3,6,10},{2,4,7},{4,7,11},{5,8,12},
		{0,2,5},{2,5,9},{5,9,14},{1,4,8},{4,8,13},{3,7,12},
		{3,4,5},{6,7,8},{7,8,9},{10,11,12},{11,12,13},{12,13,14}}

	// one token move
  var i = 0
  var j = 0
  for i = 0; i < 15; i++ {
    //  need to skip already set token
    if py.board.get(i) == 0 {
      arr[j] = int(math.Pow(2,float64(i)) )
      j ++
    }
  }

  // two token move
  for _, r := range temp2 {
    if py.board.get(r[0]) == 0 && py.board.get(r[1]) == 0  {
      arr[j] = int(math.Pow(2,float64(r[0])) + math.Pow(2,float64(r[1])))
      j = j + 1
    }
  }

  // three token move
  for _, r := range temp3 {
    if py.board.get(r[0]) == 0 && py.board.get(r[1]) == 0 && py.board.get(r[2]) == 0 {
      arr[j] = int(math.Pow(2,float64(r[0])) +
                            math.Pow(2,float64(r[1])) +
                            math.Pow(2,float64(r[2])))
      j = j + 1
    }
  }
  // slice the array
  valid_arr := arr[:j]
  sort.Ints(valid_arr)
  py.valid_moves = valid_arr
  fmt.Println("py.valid_moves=",py.valid_moves)
}


// TODO: not done yet
func (py *pyramid) my_move(){
  p := fmt.Println
  py.gen_valid_moves()
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
  pl("py.gen_valid_moves()=")
  py.gen_valid_moves()

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
          py.my_move()
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
