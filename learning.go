package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + "," + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " ,"))
// 	fmt.Println(os.Args[1:])
// }

// func main() {
// 	start := time.Now()
// 	for i := 1; i < len(os.Args); i++ {
// 		fmt.Println(i, " ", os.Args[i])
// 	}
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// }

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%\n", n, line)
		}
	}
}
