package main

type Node struct {
	NextNode *Node
}

// 1. 递归反转单链表
func reverse(headNode *Node) *Node {
	if headNode == nil {
		return headNode
	}
	if headNode.NextNode == nil {
		return headNode
	}
	var newNode = reverse(headNode.NextNode)
	headNode.NextNode.NextNode = headNode
	headNode.NextNode = nil
	return newNode
}

// 2. 快速排序: 升序
func quickAscendingSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickAscendingSort(arr, start, j)
		}
		if end > i {
			quickAscendingSort(arr, i, end)
		}
	}
}

// 题目描述：给定一个源串和目标串，能够对源串进行如下操作：
//   1.在给定位置上插入一个字符
//   2.替换任意字符
//   3.删除任意字符
// 最短编辑距离: dp[i][j]=min{dp[i−1][j]+1, dp[i][j−1]+1, dp[i−1][j−1]+(S[i]==T[j] ? 0 : 1) }
func EditDistance(src, dst string) int {
	sLen := len(src)
	dLen := len(dst)
	dp := make([][]int, sLen+1)
	dp[0] = make([]int, dLen+1)
	// 边界dp[i][0] = i，dp[0][j] = j
	for i := 1; i <= sLen; i++ {
		dp[i] = make([]int, dLen+1)
		dp[i][0] = i
	}
	for j := 1; j <= dLen; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= dLen; j++ {
			if src[i-1] == dst[j-1] {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])+1)
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j], dp[i][j-1])+1)
			}
		}
	}
	return dp[sLen][dLen]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
