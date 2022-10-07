package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		file2, _ := ioutil.ReadAll(os.Stdin)
		os.Stdout.Write(file2)

	}
	for i := 1; i < len(os.Args); i++ {
		file, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			s := "ERROR: open " + os.Args[i] + ": no such file or directory\n"
			for i := range s {
				z01.PrintRune(rune(s[i]))
			}
			os.Exit(1)
		} else {
			os.Stdout.Write(file)
		}
	}
}
