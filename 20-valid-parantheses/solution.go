package validparantheses

import "container/list"

func isValid(s string) bool {
	stk := list.New()

	for _, r := range s {
		if stk.Front() == nil {
			stk.PushFront(r)
			continue
		}

		v, _ := stk.Front().Value.(rune)
		if r == inverse[v] {
			stk.Remove(stk.Front())
			continue
		}

		stk.PushFront(r)
	}

	return stk.Len() == 0
}

var inverse = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}
