/*
	calculate
	#1 Calculate 使用栈实现四则运算计算器
	#2 Infix2Postfix 将前缀表达式转后缀表达式
*/
package calculate

import (
	"algo/internal/stack"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidInput = errors.New("invalid expression")

var np = map[rune]int{ // 运算符优先级
	'+': 1, '-': 1,
	'*': 2, '/': 2,
	'(': 3,
}

// Infix2Postfix 前缀表达式转后缀表达式
func Infix2Postfix(in string) (post string, err error) {
	var cs = &stack.Stack{} // 运算符栈
	cs.Init()

	var tmp string // 存储运算数
	var chars = []rune(in)
	for _, v := range chars {
		if (v > '0' && v < '9') || v == '.' { // 拼接运算数
			tmp += string(v)
		} else {
			// 运算数直接压入结果栈
			if tmp != "" {
				if post == "" {
					post = tmp
				} else {
					post = fmt.Sprintf("%s %s", post, tmp)
				}
				tmp = "" // 清空
			}
			switch { // 处理运算符
			// @1 若运算符栈为空，直接入栈；
			// @2 若栈顶为'('或栈顶元素优先级小于当前运算符，直接入栈；
			// @3 若栈顶运算符优先级大于等于当前运算符，则出栈运算符，压入结果栈，
			//    直至遇到小于当前运算符的栈顶，或遇到'('，或运算符栈被取空，再入栈.
			case v == '+' || v == '-' || v == '*' || v == '/':
				top, elem := cs.Top()
				// @1 @2
				if top == -1 || elem == '(' || np[v] > np[elem.(rune)] {
					if err = cs.Push(v); err != nil {
						return "", err
					}
					continue
				}
				// @3
				for np[v] <= np[elem.(rune)] {
					elem, err = cs.Pop() // 出栈栈顶运算符
					if err != nil {
						return "", err
					}
					post = fmt.Sprintf("%s %c", post, elem.(rune)) // 压入结果栈
					top, elem = cs.Top()
					// 运算符栈中的元素被取光 或 栈顶为 '('.
					if top == -1 || elem.(rune) == '(' {
						if err = cs.Push(v); err != nil {
							return "", err
						}
						break
					}
				}
			case v == '(': // 遇左括号直接入栈
				if err = cs.Push(v); err != nil {
					return "", err
				}
			case v == ')': // 右括号不入栈，并出栈运算符，直至遇到'('
				for {
					var elem any
					elem, err = cs.Pop()
					if err != nil {
						return "", err
					}
					if elem.(rune) != '(' {
						post = fmt.Sprintf("%s %c", post, elem.(rune)) // 压入结果栈
					} else { // 遇到'('，不需要将'('压入结果栈，当前运算符')'也不需要压入结果栈
						break
					}
				}
			case v == ' ': // 忽略空格
				continue
			default: // 非法字符
				return "", ErrInvalidInput
			}
		}
	}

	// 最后将剩余的运算数和运算符栈内剩余运算符压入结果栈
	if tmp != "" {
		if post == "" {
			post = tmp
		} else {
			post = fmt.Sprintf("%s %s", post, tmp)
		}
		tmp = "" // 清空
	}
	if cs.Length() != 0 {
		for {
			var elem any
			elem, err = cs.Pop()
			if err != nil {
				if errors.Is(err, stack.ErrEmpty) {
					return post, nil
				} else {
					return "", err
				}
			}
			post = fmt.Sprintf("%s %c", post, elem.(rune))
		}
	}
	return
}

// Calculate 计算器
func Calculate(in string) (ret float64, err error) {
	var postfix string
	postfix, err = Infix2Postfix(in)      // 中缀表达式转后缀表达式
	var exp = strings.Split(postfix, " ") // 按" "将后缀表达式分解为[]string

	var ns = &stack.Stack{} // 运算数栈
	ns.Init()

	// @1 字符串若为运算符，则从运算数栈中连续取出栈顶的两个运算数，执行一次计算，并将结果压入运算数栈；
	// @2 若为运算数，则直接压入运算数栈；
	for _, v := range exp {
		// @1
		if v == "+" || v == "-" || v == "*" || v == "/" {
			var elem any
			var fn, sn float64
			elem, err = ns.Pop()
			if err != nil {
				return 0, err
			}
			sn = elem.(float64)
			elem, err = ns.Pop()
			if err != nil {
				return 0, err
			}
			fn = elem.(float64)

			ret = exec(fn, sn, v)
			if err = ns.Push(ret); err != nil {
				return 0, err
			}
		} else { // @2
			var num float64
			num, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, err
			}
			if err = ns.Push(num); err != nil {
				return 0, err
			}
		}
	}
	return
}

// 执行计算
func exec(fn, sn float64, opt string) (ret float64) {
	switch opt {
	case "+":
		ret = fn + sn
	case "-":
		ret = fn - sn
	case "*":
		ret = fn * sn
	case "/":
		ret = fn / sn
	}
	return
}
