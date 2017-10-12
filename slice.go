// _Slice_是Go中一个关键的数据接口，比数组更强大的序列接口
package main

import "fmt"

func main() {
	//slice 的类型仅有它所包含的元素决定（不像
	// 数组中还需要元素的个数）。要创建一个长度非零的空
	// slice，需要使用内建的方法 `make`。这里我们创建了一
	// 个长度为3的 `string` 类型 slice（初始化为零值）
	// TODO:　更详细的slice的介绍
	s := make([]string, 3)
	fmt.Println("emp:", s) //emp: [  ]

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s: ", s) //s:  [a b c]

	// slice支持比数组更多的操作 比如append
	s = append(s, "b")
	s = append(s, "e", "f")
	fmt.Println("append: ", s) //append:  [a b c b e f]

	// slice可以被copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("C copy: ", c) //C copy:  [a b c b e f]

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操
	// 作。例如，这里得到一个包含元素 `s[2]`, `s[3]`,
	// `s[4]` 的 slice。

	l := s[2:5]
	fmt.Println("slice1: ", l) //slice1:  [c b e]
	// 这个 slice 从 `s[0]` 到（但是不包含）`s[5]`。
	l = s[:5]
	fmt.Println("sl2:", l) //sl2: [a b c b e]

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不
	// 同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) //2d:  [[0] [1 2] [2 3 4]]
}
