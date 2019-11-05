package algo

type Node struct {
	NextNode *Node
}

// 递归反转单链表
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

// 快速排序: 升序
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
