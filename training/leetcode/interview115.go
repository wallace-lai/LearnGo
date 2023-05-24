// 二分查找

package main

// v1 : 4ms
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	// 先在第1列中进行二分查找
	left := 0
	right := m - 1
	for left <= right {
		mid := left + (right - left) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else if matrix[mid][0] > target {
			right = mid - 1
		}
	}

	// 若target比matrix[0][0]还小
	if right == -1 {
		return false
	}

	// 再在第row行进行二分查找
	row := right
	left = 0
	right = n - 1
	for left <= right {
		mid := left + (right - left) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else if matrix[row][mid] > target {
			right = mid - 1
		}
	}

	return false
}

