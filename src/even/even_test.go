// even test
package even

//import "testing"
import "fmt"

/*
func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even!~") //用默认格式对参数格式化
		t.Fail()                    //标记函数失败，但仍继续执行。FailNow立即中断
	}
}
*/
func ExampleEven() {
	if even(3) {
		fmt.Printf("Is odd\n")
	}
	// Output:
	// Is odd
}
