package bit

import (
	"fmt"
	"testing"
)

var res int

func BenchmarkGetMin1(b *testing.B) {
	fmt.Println("-------------")
	fmt.Println(GetMin1(5, 2))
	fmt.Println(GetMin2(5, 2))
	fmt.Println("-------------")
	var t int
	for n := 0; n < b.N; n++ {
		//	t = GetMin1(5, 2)
	}
	res = t
}

func BenchmarkGetMin2(b *testing.B) {
	var t int
	for n := 0; n < b.N; n++ {
		//	t = GetMin1(5, 2)
	}
	res = t
}
