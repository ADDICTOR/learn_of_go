// echo3 输出其命令行参数os.Arg[0]
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
