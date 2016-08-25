// stack
package stack

//	"fmt"

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(x int) {
	if s.i >= 10 {
		return
	}
	s.data[s.i] = x
	s.i++
}

func (s *Stack) Pop() int {
	if s.i < 1 {
		return -9999
	}
	s.i--
	return (s.data[s.i])
}
