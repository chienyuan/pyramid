package main

import (
  "fmt"
  "bufio"
  "os"
  "math"
  "math/rand"
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

type game struct {
  steps[] int
}

// token val 
// start = 000000000000000 
// 0 : free
// 1 : non-free
type board struct {
  val int
}

type pyramid struct {
  board board
  root,temp,node_buff  node
  win[6299] uint
  valid_moves[] int
}

//     val= 111000000000
//  pos 2 = 010000000000 and
//    ret   010000000000 
// so if pos == ret , return 1
func (b board) get(pos int) int {
  v := int(math.Pow(2,float64(pos)))
  if ( v & b.val )  == v {
    return 1
  }
  return 0
}

func (b *board) set(pos int) {
  v := int(math.Pow(2,float64(pos)))
  b.val |= v
}

func (b *board) addvals(move int) {
  b.val += move
}

func (b board) vals() int {
  return b.val
}

func (b *board) reset() {
  b.val = 0
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
  pl( py.board.vals())
}

func (py *pyramid ) human_move()  int{
  pl := fmt.Println
  p  := fmt.Print
  py.gen_valid_moves()
  p("Please enter your move (A-N)? ")
  reader := bufio.NewReader(os.Stdin)
  var str string
  var move int
  for {
    str,_ = reader.ReadString('\n')
    str = str[:len(str)-1]

    pl("len(str)=",len(str))
    // only allow pick 1,2,3 token
    var valid bool
    move,valid =  py.is_valid_move(str)
    if valid {
      break
    }

    p("Invalid !!! Please enter your move again (A-N)? ")
  }

  for _, r := range str {
    py.board.set( int(r - 65) )
  }

  py.display()

  return move
}

func (py pyramid) is_valid_move ( str string) ( int,bool ) {
    if len(str) < 0 || len(str) > 3 { 
      return 0,false
    }

    var move int
    for _, r := range str {
      move += int( math.Pow(2,float64(r - 65) ) )
    }
    fmt.Println("move=",move)

    for _,r := range py.valid_moves {
      if move == r {
        return move ,true
      }
    }

  return move,false
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


func (py *pyramid) compute_move(play bool) int{
  p := fmt.Println
  py.gen_valid_moves()

  // TODO: add AI logic here
  
  pick := rand.Intn(len(py.valid_moves))
  var move int = py.valid_moves[pick]
  if play {
    p("len(valid_moves)",len(py.valid_moves))
    p("my move valid_moves[", pick ,"]=",move)
  }
  py.board.addvals(move)
  py.display()
  return move
}

func (py pyramid) is_gameover() bool {
  for i:=0 ; i < 15 ; i++ {
    if py.board.get(i) == 0 {
      return false
    }
  }
  return true
}


// TODO: need to add game state, step , win strut
func gen() (map[int]int,bool) {

  steps := make(map[int]int)
  py := pyramid{}
  s := 0
  win := false

  for !py.is_gameover() {
    steps[s]= py.compute_move(false);
    s++
    if py.is_gameover() {
      win = true
     } else {
      steps[s]=py.compute_move(false)
      s++
      if py.is_gameover() {
        win = false
      }
    }
  }
  return steps,win
}

func main(){

  pl := fmt.Println
  games := make(map[int] map[int]int)
  var win bool
  for i:=0 ; i < 100 ; i++ {
    games[i],win = gen()
  }
  pl("games: win=",win," steps",games)

  //play()
}

func play(){
  pl := fmt.Println

  py := pyramid{}
  pl("py.gen_valid_moves()=")
  py.gen_valid_moves()

	pl("Welcome to Pyramid game ")
	pl("The Rules of Pyramid: Players alternate take out one to three ")
	pl("tokens in one line. Who takes the last one token is loser")

  for {
    py.board.reset()
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Do you want to move first(y/n)?")
    ans,_ := reader.ReadByte()
    _,_ = reader.ReadByte()
    s     := 0
    steps := make(map[int]int)

    if ans == 'y'  {
      py.display()
      pl("OK ! You move first.")
      for !py.is_gameover() {
        steps[s]= py.human_move();
        s++
        if py.is_gameover() {
          pl("I win")
        } else {
          steps[s]=py.compute_move(true)
          s++
          if py.is_gameover() {
            pl("You win")
          }
        }
      }
    } else {
      pl(" OK ! I  move first.")
      for !py.is_gameover() {
        steps[s]=py.compute_move(true)
        s++
        if py.is_gameover() {
          pl("You win")
        } else  {
          steps[s]=py.human_move()
          s++
          if py.is_gameover() {
            pl("I win")
          }
        }
      }
    }

    pl("Game Over")
    pl("Game Result:", steps)
    pl("Do you want to play again(y/n)?")
    ans,_ = reader.ReadByte()
    _,_ = reader.ReadByte()

    if  ans == 'n'   { break }
  }

}
