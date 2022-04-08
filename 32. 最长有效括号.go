package hot100

/*
始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，
这样的做法主要是考虑了边界条件的处理，栈里其他元素维护左括号的下标：

对于遇到的每个 ‘(’，将它的下标放入栈中
对于遇到的每个 ‘)’，先弹出栈顶元素表示匹配了当前右括号：

如果栈为空，说明当前的右括号为没有被匹配的右括号，将其下标放入栈中来更新之前提到的
「最后一个没有被匹配的右括号的下标」
如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」

从前往后遍历字符串并更新答案即可。
需要注意的是，如果一开始栈为空，第一个字符为左括号的时候会将其放入栈中，
这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，在一开始的时候往栈中放入 −1。

*/
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	ans := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if ans < i-stack[len(stack)-1] {
					ans = i - stack[len(stack)-1]
				}
			}
		}
	}
	return ans
}
