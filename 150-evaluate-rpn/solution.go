package evaluaterpn

import (
	"container/list"
	"strconv"
)

func EvalRPN(tokens []string) int {
	return evalRPN(tokens)
}

func evalRPN(tokens []string) int {
	stk := list.New()

	for _, token := range tokens {
		op, ok := operators[token]
		if !ok {
			stk.PushFront(mustAtoI(token))
			continue
		}
		opd2, opd1 := stk.Remove(stk.Front()).(int), stk.Remove(stk.Front()).(int)
		stk.PushFront(op.Operate(opd1, opd2))
	}

	return stk.Remove(stk.Front()).(int)
}

func mustAtoI(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var operators = map[string]Operator{
	"+": OperatorFunc(func(x, y int) int {
		return x + y
	}),
	"-": OperatorFunc(func(x, y int) int {
		return x - y
	}),
	"/": OperatorFunc(func(x, y int) int {
		return x / y
	}),
	"*": OperatorFunc(func(x, y int) int {
		return x * y
	}),
}

type Operator interface {
	Operate(x, y int) int
}

type OperatorFunc func(x, y int) int

func (o OperatorFunc) Operate(x, y int) int {
	return o(x, y)
}
