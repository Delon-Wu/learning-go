package task

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
var countMap = make(map[int]int)

func SingleNumber(nums []int) int {
	for i, _ := range nums {
		countMap[nums[i]]++
	}
	var target int
	for k, v := range countMap {
		if v == 1 {
			target = k
			break
		}
	}
	return target
}

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。

func IsPalindrome(x int) bool {
	var newNum = 0
	for x > 0 {
		newNum = newNum*10 + x%10
	}
	return newNum == x
}

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

func IsValid(s string) bool {
	var stack []string
	var matchMap = map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for letter := range s {
		letterString := string(s[letter])
		if len(stack) > 0 && stack[len(stack)-1] == matchMap[letterString] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, letterString)
		}
	}
	return len(stack) == 0
}

//最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。

//func LongestCommonPrefix(strs []string) string {
//	prefix := strs[0][0:1]
//	for _, str := range strs {
//		if
//	}
//}
