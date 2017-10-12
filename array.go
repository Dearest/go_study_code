package main

import "fmt"

func main() {
	var a [5]int           // 定义一个长度为5的数组，数组默认是零值。
	fmt.Println("emp:", a) // 输出为 emp: [0,0,0,0,0]

	a[4] = 100 // 设置指定位置的值
	fmt.Println("set: ", a)
	fmt.Println("get", a[4])
	fmt.Println("len", len(a)) //内置函数返回数组的长度 用法跟c python相似

	b := [5]int{1, 2, 3, 4, 5} //初始化数组
	fmt.Println("dcl", b)

	//多维数组
	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD:", twoD) //twoD: [[0 1 2] [1 2 3]]

}
