package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type mode int

const (
	noneMode mode = iota
	addMode
	subMode
	mulMode
	divMode
)

func performOperation(m mode, r1, r2 int) int {
	var a int
	switch m {
	case addMode:
		a = r1 + r2
	case subMode:
		a = r1 - r2
	case mulMode:
		a = r1 * r2
	case divMode:
		a = r1 / r2
	case noneMode:
		fallthrough
	default:
		fmt.Printf("Warning: impossible operator mode entered: %d\n", m)
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	mode := noneMode
	r1 := 0 // register for mathematical operations
	r2 := 0

	operator := regexp.MustCompile("[\\-+*/]")

	fmt.Println("Type an equation to solve")

	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	for _, c := range line {
		s := string(c)
		isNumber := unicode.IsNumber(c)
		isOperator := operator.MatchString(s)

		if isNumber {
			n, _ := strconv.Atoi(s)
			if mode == noneMode {
				r1 *= 10
				r1 += n
			} else {
				r2 *= 10
				r2 += n
			}
		} else if isOperator {
			if r2 != 0 {
				r1 = performOperation(mode, r1, r2)
				r2 = 0
			}
			switch s {
			case "+":
				mode = addMode
			case "-":
				mode = subMode
			case "*":
				mode = mulMode
			case "/":
				mode = divMode
			default:
				fmt.Printf("Warning: unknown operator entered: %s\n", s)
			}
		}
	}

	r1 = performOperation(mode, r1, r2)

	fmt.Printf("%d", r1)
}
