package piscine

import (
	"fmt"
	"os"
)

func PrintNbr(n int) {
	fmt.Println(len(os.Args), os.Args)
}
