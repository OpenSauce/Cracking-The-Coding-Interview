package chapter_five

import "fmt"

/*
Bit Manipulation By Hand

Column 1
0110 + 0010 = 1000 /
0011 + 0010 = 0101 /
0110 - 0011 = 0011 /
1000 - 0110 = 0010 /

Column 2
0011 * 0101 = 1111 /
0011 * 0011 = 1001 /
1101 >> 2   = 0110 X 0011
1101 ^ 0101 = 1000 /

Column 3
0110 + 0110 = 1100 /
0100 * 0011 = 1100 /
1101 ^ (~1101) = 1101 ^ 0010 = 1111 /
1011 & (~0 << 2) = 1011 & 1111 << 2 = 1010 X 1000
*/

func QuestionOne(n, m int32, j, i int32) int32 {
	if i < 0 || i > j || j >= 32 {
		return 0
	}

	allOne := ^0

	var left int
	if j < 31 {
		left = allOne<<j + 1
	} else {
		left = 0
	}

	right := ((1 << i) - 1)

	mask := left | right

	clrN := n & int32(mask)
	shiftM := m << i
	return clrN | shiftM
}

func QuestionTwo(number float32) string {
	if number >= 1 || number <= 0 {
		return "ERROR"
	}

	result := "."
	fraction := float32(0.5)
	for number > 0 {
		if len(result) > 32 {
			return "ERROR"
		}

		if number >= fraction {
			result += "1"
			number -= fraction
		} else {
			result += "0"
		}
		fraction /= 2
	}
	return result
}

func QuestionThree(num int32) int {
	if ^num == 0 {
		return 0
	}

	currentLength := 0
	previousLength := 0
	maxLength := 1
	for num != 0 {
		if (num & 1) == 1 {
			currentLength++
		} else if (num & 1) == 0 {
			if num&2 == 0 {
				previousLength = 0
			} else {
				previousLength = currentLength
			}
			currentLength = 0
		}
		if (previousLength + currentLength + 1) > maxLength {
			maxLength = previousLength + currentLength + 1
		}
		num >>= 1
	}
	return maxLength
}

// func QuestionFour(num int32) {
// 	var found1, found2 bool
// 	i := 0
// 	n := num
// 	for n != 0 {
// 		if found1 && found2 {
// 			return
// 		}
// 		if (num&(1<<i)) != 0 && !found1 {
// 			if (num & (1 << (i + 1))) == 0 {
// 				temp := num & ^(1 << i)
// 				temp |= (1 << (i + 1))
// 				fmt.Println(temp)
// 				found1 = true
// 				continue
// 			}
// 		} else if (num&(1<<i)) == 0 && !found2 {
// 			if (num & (1 << (i + 1))) != 0 {
// 				temp := num & ^(1 << (i + 1))
// 				temp |= (1 << i)
// 				fmt.Println(temp)
// 				found2 = true
// 				continue
// 			}
// 		}
// 		i++
// 		n >>= 1
// 	}
// }

func QuestionFour(num int) {
	fmt.Println(getNext(num))
	fmt.Println(getPrev(num))
}

func getNext(n int) int {
	c := n
	c0 := 0
	c1 := 0

	for ((c & 1) == 0) && (c != 0) {
		c0++
		c >>= 1
	}

	for (c & 1) == 1 {
		c1++
		c >>= 1
	}

	if c0+c1 == 31 || c0+c1 == 0 {
		return -1
	}

	p := c0 + c1
	n |= (1 << p)
	n &= ^((1 << p) - 1)
	n |= (1 << (c1 - 1)) - 1
	return n
}

func getPrev(n int) int {
	temp := n
	c0 := 0
	c1 := 0
	for temp&1 == 1 {
		c1++
		temp >>= 1
	}

	if temp == 0 {
		return -1
	}

	for temp&1 == 0 && temp != 0 {
		c0++
		temp >>= 1
	}

	p := c0 + c1
	n &= (^(0) << (p + 1))

	mask := (1 << (c1 + 1)) - 1
	n |= mask << (c0 - 1)
	return n
}

/*
Question Five

Works out if it's a power of 2

Absolutely balling!

*/

func QuestionSix(a, b int) int {
	count := 0
	c := a ^ b

	for c != 0 {
		count += c & 1
		c >>= 1
	}
	return count
}

func QuestionSeven(in int) int {
	mask := 1
	for i := 0; i < 16; i++ {
		mask <<= 2
		mask++
	}

	return ((in & mask) << 1) | ((in & (mask << 1)) >> 1)
}

func QuestionEight() {

}
