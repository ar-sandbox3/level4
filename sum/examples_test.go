package sum_test

import (
	"fmt"

	"github.com/ar-sandbox3/level4/sum"
)

func ExampleInts() {
	fmt.Println(sum.Ints([]int{1, 2, 3, 4, 5}...))
	// Output:
	// 15
}
