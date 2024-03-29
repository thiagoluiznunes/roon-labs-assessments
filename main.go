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

var stack []float64
var isInterativeMode = false

func printSlice(s []float64) {
	if len(s) == 0 {
		return
	}
	value := fmt.Sprintf("%v", s[0])
	if value != "0" {
		if !isInterativeMode {
			fmt.Println(value)
		} else {
			fmt.Printf("%v ", value)
		}

	} else if value == "0" && len(stack) == 1 {
		if !isInterativeMode {
			fmt.Println(value)
		} else {
			fmt.Printf("%v ", value)
		}
	}
	printSlice(s[1:])
}

func execOperationInStack(operation string) {
	var result float64
	for index, value := range stack {
		switch operation {
		case "+":
			result = value + result
		case "-":
			if index != 0 {
				result = result - value
				continue
			}
			result = value - result
		case "/":
			result = value / result
		case "*":
			if index != 0 {
				result = value * result
				continue
			}
			result = value
		}
	}
	stack = []float64{result}
}

func checkIsAllowedExpression(tks []string) bool {
	expression := strings.Join(tks[:], " ")
	for _, value := range domain.NOT_ALLOWED_EXPRESSIONS {
		if ok := strings.Contains(expression, value); ok {
			return false
		}
	}

	return true
}

func evalrpn(tks []string, interactiveMode bool) {
	var isZero bool
	var x float64

	if ok := checkIsAllowedExpression(tks); !ok {
		if interactiveMode {
			printSlice(stack)
			fmt.Printf("> ")
		}
		return
	}
	pop := func() float64 {
		n := len(stack) - 1
		t := stack[n]
		stack = stack[:n]
		return t
	}
	defer func() {
		if recover() == nil {
			if interactiveMode {
				printSlice(stack)
				fmt.Printf("> ")
			} else {

				printSlice(stack)
			}
		} else {
			if interactiveMode {
				printSlice(stack)
				fmt.Printf("> ")
			} else {
				printSlice(stack)
			}
		}
	}()
	executed := false
	for _, tk := range tks {
		if tk == "pi" {
			tk = fmt.Sprintf("%v", domain.PI)
		}
		if tk == "q" {
			gracefulShutdown()
		}
		switch tk {
		case "+":
			if len(tks) == 1 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(tks) > 3 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(stack) > 1 {
				var pop1 float64
				if isZero {
					pop1 = 0
				} else {
					pop1 = pop()
				}
				pop2 := pop()
				x = pop2 + pop1
			}
		case "*":
			if len(tks) == 1 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(tks) > 3 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(tks) > 1 {
				var pop1 float64
				if isZero {
					pop()
					break
				} else {
					pop1 = pop()
				}
				pop2 := pop()
				x = pop2 * pop1
			}
		case "-":
			if len(tks) == 1 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(tks) > 3 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(stack) > 1 {
				var pop1 float64
				if isZero {
					pop1 = 0
				} else {
					pop1 = pop()
				}
				pop2 := pop()
				x = pop2 - pop1
			}
		case "/":
			if len(tks) == 1 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(tks) > 3 {
				execOperationInStack(tk)
				executed = true
				break
			}
			if len(stack) > 1 {
				var pop1 float64
				if isZero {
					pop()
					break
				} else {
					pop1 = pop()
				}
				pop2 := pop()
				x = pop2 / pop1
			}
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
		if x != 0 && !executed {
			stack = append(stack, x)
		} else if len(stack) == 0 {
			stack = append(stack, x)
		} else {
			isZero = true
		}
	}
}

func gracefulShutdown() {
	fmt.Println(domain.GRACEFUL_SHUTDOWN_TEXT)
	os.Exit(1)
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
				evalrpn(tks, isInterativeMode)
			}
		}
	} else {
		isInterativeMode = true
		stdin := bufio.NewReader(os.Stdin)
		fmt.Printf("> ")
		for {
			s, e := stdin.ReadString('\n')
			if e != nil {
				break
			}
			if tks := strings.Fields(s); len(tks) > 0 {
				evalrpn(tks, isInterativeMode)
			}
		}
	}
}
