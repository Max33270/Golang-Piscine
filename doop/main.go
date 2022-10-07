package main

import "os"

func Atoi(s string) int {
	if s == "" {
		os.Exit(0)
	}
	a := []rune(s)
	c := 0
	for b, v := range s {
		d := int(a[b])
		v++
		if (d == 45 || d == 43) && b == 0 {
		} else if !(d <= 57 && d >= 48) {
			os.Exit(0)
		} else {
			c = c*10 + d - 48
		}
	}
	if s[0] == 45 {
		return -c
	}
	return c
}

func main() {
	tab := []byte{}
	if len(os.Args) == 4 {
		if len(os.Args[1]) < 19 && len(os.Args[3]) < 19 {
			x := Atoi(os.Args[1])
			y := Atoi(os.Args[3])
			if os.Args[2] == "+" {
				for result := x + y; result != 0; result = result / 10 {
					if result < 0 {
						os.Stdout.WriteString("-")
						result = result * (-1)
					}
					a := result % 10
					tab = append(tab, byte(a)+48)
				}
			} else if os.Args[2] == "-" {
				for result := x - y; result != 0; result = result / 10 {
					if result < 0 {
						os.Stdout.WriteString("-")
						result = result * (-1)
					}
					a := result % 10
					tab = append(tab, byte(a)+48)
				}
			} else if os.Args[2] == "*" {
				for result := x * y; result != 0; result = result / 10 {
					if result < 0 {
						os.Stdout.WriteString("-")
						result = result * (-1)
					}
					a := result % 10
					tab = append(tab, byte(a)+48)
				}
			} else if os.Args[2] == "/" {
				if y == 0 {
					os.Stdout.WriteString("No division by 0")
				} else {
					for result := x / y; result != 0; result = result / 10 {
						if result < 0 {
							os.Stdout.WriteString("-")
							result = result * (-1)
						}
						a := result % 10
						tab = append(tab, byte(a)+48)
					}
				}
			} else if os.Args[2] == "%" {
				if y == 0 {
					os.Stdout.WriteString("No modulo by 0")
				} else {
					for result := x % y; result != 0; result = result / 10 {
						if result < 0 {
							os.Stdout.WriteString("-")
							result = result * (-1)
						}
						a := result % 10
						tab = append(tab, byte(a)+48)
					}
				}
			} else {
				return
			}
			for i := len(tab) - 1; i >= 0; i-- {
				os.Stdout.WriteString(string(tab[i]))
			}
			os.Stdout.WriteString("\n")
		}
	}
}
