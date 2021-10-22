package chapter_five

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
	return ""
}
