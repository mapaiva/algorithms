package brackets

import "container/list"

// IsBalanced return wether a given string has a balanced brackets relation.
func IsBalanced(s string) bool {
	stack := list.New()
	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack.PushFront(char)
			continue
		}

		if stack.Len() == 0 {
			return false
		}

		switch char {
		case '}':
			el := stack.Front()
			stack.Remove(el)

			if el.Value.(rune) != '{' {
				return false
			}
		case ']':
			el := stack.Front()
			stack.Remove(el)

			if el.Value.(rune) != '[' {
				return false
			}
		case ')':
			el := stack.Front()
			stack.Remove(el)

			if el.Value.(rune) != '(' {
				return false
			}
		}
	}

	return stack.Len() == 0
}
