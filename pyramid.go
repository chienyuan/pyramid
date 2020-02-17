package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	//"github.com/goml/gobrain"
)

type node struct {
	step      uint
	bestchild uint
	children  uint
	cnum      int
	value     int
}

type game struct {
	steps []int
}

// start = 000000000000000
// 0 : free
// 1 : non-free
type board struct {
	state int
}

type pyramid struct {
	board board
	level int
	win   [6299]uint
	//validMoves[] int
}

// Checking the state of  position
// state = 111000000000
// pos 2 = 010000000000 and
//   ret   010000000000
// so if pos == ret , return 1
func (b board) isTaken(pos int) bool {
	v := int(math.Pow(2, float64(pos)))
	if (v & b.state) == v {
		return true
	}
	return false
}

func (b board) isFree(pos int) bool {
	v := int(math.Pow(2, float64(pos)))
	if (v & b.state) == v {
		return false
	}
	return true
}

func (b *board) set(pos int) {
	v := int(math.Pow(2, float64(pos)))
	b.state |= v
}

func (b *board) unset(pos int) {
	v := int(math.Pow(2, float64(pos)))
	b.state ^= v
}

func (b *board) placeMove(move int) {
	b.state += move
}

func (b *board) removeMove(move int) {
	b.state -= move
}

func (b board) getState() int {
	return b.state
}

// after play pos , it only got one token left?
func (b board) isLastMove(move int) bool {
	b.placeMove(move)
	var last bool = false

	if b.state == 32767 {
		last = true
	}
	b.removeMove(move)
	return last
}

func (b *board) reset() {
	b.state = 0
}

//  0(0,4)
//  1(1,3) 2(1,5)
//  3(2,2) 4(2,4) 5(2,6)
//  6(3,1) 7(3,3) 8(3,5) 9(3,7)
//  A(4,0) B(4,2) C(4,4) D(4,6) E(4,8)
func (py pyramid) display() {
	d := 5 // depth
	t := 0 // start token value
	for i := 0; i < d; i++ {
		// print prefix space
		for k := 1; k < d-i; k++ {
			fmt.Print(" ")
		}

		// print token
		for j := 0; j < i+1; j++ {
			if py.board.isTaken(t) {
				fmt.Print("*")
			} else {
				fmt.Print(string(t + 65))
			}
			fmt.Print(" ")
			t++
		}

		// print endl
		fmt.Println("")
	}
	fmt.Println(py.board.getState())
}

func (py *pyramid) humanMove() int {
	p := fmt.Print
	var validMoves []int = py.board.getValidMoves()
	p("Please enter your move (A-N)? ")
	reader := bufio.NewReader(os.Stdin)
	var str string
	var move int
	for {
		str, _ = reader.ReadString('\n')
		str = str[:len(str)-1]

		var valid bool
		move, valid = isValidMove(str, validMoves)
		if valid {
			break
		}

		p("Invalid !!! Please enter your move again (A-N)? ")
	}

	for _, r := range str {
		py.board.set(int(r - 65))
	}

	py.display()

	return move
}

func max(v1 int, v2 int) int {
	if v2 >= v1 {
		return v2
	} else {
		return v1
	}
}
func min(v1 int, v2 int) int {
	if v2 <= v1 {
		return v2
	} else {
		return v1
	}
}

func (b board) minimax(move int, depth int, maxPlay bool) int {
	last := b.isLastMove(move)

	if last {
		if maxPlay {
			return 1
		}
		return -1
	}

	if depth == 0 {
		return 0
	}

	b.placeMove(move)
	var validMoves []int = b.getValidMoves()

	var value int = 0
	if maxPlay {
		value = -1
		for _, child := range validMoves {
			v := b.minimax(child, depth-1, false)
			value = max(value, v)
		}
	} else {
		value = 1
		for _, child := range validMoves {
			v := b.minimax(child, depth-1, true)
			value = min(value, v)
		}
	}
	b.removeMove(move)
	return value
}

func (py *pyramid) computeMove(play bool) int {
	var cb board
	cb.state = py.board.getState()

	var validMoves []int = cb.getValidMoves()
	validLen := len(validMoves)
	fmt.Print("len=", validLen, ":")

	// TODO: add AI logic here
	var pick int = -1
	for i, move := range validMoves {
		fmt.Print(".")
		r := cb.minimax(move, py.level, false)
		if r > 0 {
			pick = i
			fmt.Println("winning pick:", pick, " validMoves[]=", validMoves, " break ")
			break
		}
	}

	// random move
	if pick == -1 {
		pick = rand.Intn(len(validMoves))
		fmt.Println("random  pick:", pick, " validMoves[]=", validMoves)
	}

	var move int = validMoves[pick]

	if play {
		//p("len(validMoves)", len(validMoves))
		//p("my move validMoves[", pick, "]=", move)
	}
	py.board.placeMove(move)
	py.display()
	return move
}

func isValidMove(str string, validMoves []int) (int, bool) {
	if len(str) < 0 || len(str) > 3 {
		return 0, false
	}

	var move int
	// convert string to integer move value
	for _, r := range str {
		move += int(math.Pow(2, float64(r-65)))
	}

	// check if the value in validMoves
	for _, r := range validMoves {
		if move == r {
			return move, true
		}
	}

	return move, false
}

func (b board) getValidMoves() []int {
	//    0
	//   1 2
	//  3 4 5
	// 6 7 8 9
	//a b c d e
	// two token move data
	var arr [63]int
	var temp2 = [30][2]int{
		{0, 1}, {1, 3}, {3, 6}, {6, 10}, {2, 4}, {4, 7}, {7, 11}, {5, 8}, {8, 12}, {9, 13},
		{0, 2}, {2, 5}, {5, 9}, {9, 14}, {1, 4}, {4, 8}, {8, 13}, {3, 7}, {7, 12}, {6, 11},
		{1, 2}, {3, 4}, {4, 5}, {6, 7}, {7, 8}, {8, 9}, {10, 11}, {11, 12}, {12, 13}, {13, 14}}

	// three token move data
	var temp3 = [18][3]int{
		{0, 1, 3}, {1, 3, 6}, {3, 6, 10}, {2, 4, 7}, {4, 7, 11}, {5, 8, 12},
		{0, 2, 5}, {2, 5, 9}, {5, 9, 14}, {1, 4, 8}, {4, 8, 13}, {3, 7, 12},
		{3, 4, 5}, {6, 7, 8}, {7, 8, 9}, {10, 11, 12}, {11, 12, 13}, {12, 13, 14}}

	// one token move
	var i = 0
	var j = 0
	for i = 0; i < 15; i++ {
		//  need to skip already set token
		if b.isFree(i) {
			arr[j] = int(math.Pow(2, float64(i)))
			j++
		}
	}

	// two token move
	for _, r := range temp2 {
		if b.isFree(r[0]) && b.isFree(r[1]) {
			arr[j] = int(math.Pow(2, float64(r[0])) + math.Pow(2, float64(r[1])))
			j = j + 1
		}
	}

	// three token move
	for _, r := range temp3 {
		if b.isFree(r[0]) && b.isFree(r[1]) && b.isFree(r[2]) {
			arr[j] = int(math.Pow(2, float64(r[0])) +
				math.Pow(2, float64(r[1])) +
				math.Pow(2, float64(r[2])))
			j = j + 1
		}
	}
	// slice the array so it could sort
	validArr := arr[:j]
	sort.Ints(validArr)
	var validMoves []int = validArr
	return validMoves
}

func (py pyramid) isGameOver() bool {
	for i := 0; i < 15; i++ {
		if py.board.isFree(i) {
			return false
		}
	}
	return true
}

// TODO: need to add game state, step , win strut
func gen() (map[int]int, bool) {

	steps := make(map[int]int)
	py := pyramid{level: 3}
	s := 0
	win := false

	for !py.isGameOver() {
		steps[s] = py.computeMove(false)
		s++
		if py.isGameOver() {
			win = true
		} else {
			steps[s] = py.computeMove(false)
			s++
			if py.isGameOver() {
				win = false
			}
		}
	}
	return steps, win
}

func main() {
	var train bool

	if train {
		games := make(map[int]map[int]int)
		var win bool
		for i := 0; i < 100; i++ {
			games[i], win = gen()
		}
		fmt.Println("games: win=", win, " steps", games)
	}

	play()
}

func play() {
	pl := fmt.Println
	p := fmt.Print

	py := pyramid{level: 3}

	pl("Welcome to Pyramid game ")
	pl("The Rules of Pyramid: Players alternate take out one to three ")
	pl("tokens in one line. Who takes the last one token is loser")

	for {
		py.board.reset()
		// For Testing
		//py.board.state = 0B0111111111110000
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to move first(y/n)?")
		ans, _ := reader.ReadByte()
		_, _ = reader.ReadByte()
		s := 0
		steps := make(map[int]int)

		if ans == 'y' {
			py.display()
			pl("OK ! You move first.")
			for !py.isGameOver() {
				steps[s] = py.humanMove()
				s++
				if py.isGameOver() {
					pl("You lose,I win!!!")
				} else {
					steps[s] = py.computeMove(true)
					s++
					if py.isGameOver() {
						pl("You win")
					}
				}
			}
		} else {
			pl(" OK ! I  move first.")
			for !py.isGameOver() {
				steps[s] = py.computeMove(true)
				s++
				if py.isGameOver() {
					pl("You win")
				} else {
					steps[s] = py.humanMove()
					s++
					if py.isGameOver() {
						pl("You lose,I win!!")
					}
				}
			}
		}

		pl("Game Over")
		pl("Game Result:", steps)
		p("Do you want to play again(y/n)?")
		ans, _ = reader.ReadByte()
		_, _ = reader.ReadByte()

		if ans == 'n' {
			break
		}
	}
}
