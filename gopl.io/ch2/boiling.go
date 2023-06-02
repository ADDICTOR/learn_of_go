// boiling 输出水的沸点
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %G°C\n", f, c)
	// 输出
	// boiling point = 212°F or 100°C
}
