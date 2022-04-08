package hot100

// 栈 模拟，如果栈顶元素为右括号，则返回false
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if v == ')' {
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
					continue
				} else {
					return false
				}
			} else if v == ']' {
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
					continue
				} else {
					return false
				}
			} else if v == '}' {
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
					continue
				} else {
					return false
				}
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
