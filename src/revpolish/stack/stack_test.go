// stack
package stack

import "fmt"
import "testing"

func TestStack(t *testing.T) {
	var s Stack
	s.Push(24)
	s.Push(23)
	s.Push(2333)
	fmt.Printf("total number is: %d\n", s.i)
	fmt.Printf("2nd number is: %d\n", s.data[1])
	fmt.Printf("pop 1 is: %d\n", s.Pop())
	fmt.Printf("pop 2 is: %d\n", s.Pop())
}
