package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"roon-lab-assessments/domain"
	"strconv"
	"strings"
)

func gracefulShutdown() {
	fmt.Println(domain.GRACEFUL_SHUTDOWN_TEXT)
	os.Exit(1)
}

func evalrpn(tks []string) {
	var (
		nums []float64
		x    float64
	)
	pop := func() float64 {
		n := len(nums) - 1
		t := nums[n]
		nums = nums[:n]
		return t
	}
	defer func() {
		if recover() == nil && len(nums) == 1 {
			fmt.Printf("%v > ", x)
		} else {
			fmt.Printf("not allowed operation > ")
		}
	}()
	for _, tk := range tks {
		if tk == "pi" {
			tk = fmt.Sprintf("%v", domain.PI)
		}
		if tk == "q" {
			gracefulShutdown()
		}
		switch tk {
		case "+":
			x = pop() + pop()
		case "*":
			x = pop() * pop()
		case "-":
			x = pop()
			x = pop() - x
		case "/":
			x = pop()
			x = pop() / x
		case "sqrt":
			x = math.Sqrt(pop())
		case "cos":
			x = math.Cos(pop())
		case "sin":
			x = math.Sin(pop())
		default:
			var e error
			x, e = strconv.ParseFloat(tk, 64)
			if e != nil {
				panic(0)
			}
		}
		nums = append(nums, x)
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "-h":
			fmt.Println(domain.HELP_INFO_TEXT)
		case "--help":
			fmt.Println(domain.HELP_INFO_TEXT)
		default:
			operation := strings.Join(args[1:], " ")
			if tks := strings.Fields(operation); len(tks) > 0 {
				evalrpn(tks)
			}
			fmt.Println()
		}
	} else {
		stdin := bufio.NewReader(os.Stdin)
		fmt.Printf("> ")
		for {
			s, e := stdin.ReadString('\n')
			if e != nil {
				break
			}
			if tks := strings.Fields(s); len(tks) > 0 {
				evalrpn(tks)
			}
		}
	}
}
