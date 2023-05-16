// 回溯
package main

// https://leetcode.cn/problems/generate-parentheses/?envType=study-plan-v2&id=top-interview-150

func depthFirstSearch(leftNum int, rightNum int, str string, result *[]string) {
    if leftNum == 0 && rightNum == 0 {
        *result = append(*result, str)
        return
    }

    if leftNum > 0 {
        depthFirstSearch(leftNum - 1, rightNum, str + "(", result)
    }
    if rightNum > 0 && rightNum > leftNum {
        depthFirstSearch(leftNum, rightNum - 1, str + ")", result)
    }
}

func generateParenthesis(n int) []string {
    result := []string {}
    depthFirstSearch(n, n, "", &result)
    return result
}