package main

import "fmt"
import "time"
func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	  // 在一个 `case` 语句中，你可以使用逗号来分隔多个表
    // 达式。在这个例子中，我们很好的使用了可选的
    // `default` 分支。
	t := time.Now()
	switch {
		case t.Hour() < 12:
		fmt.Println("it`s before noon")
	default:
		fmt.Println("it's after noon")
	}
}