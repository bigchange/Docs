/*
 * @Author: Jerry You
 * @CreatedDate: 2019-11-05 11:29:30
 * @Last Modified by: Jerry You
 * @Last Modified time: 2019-11-05 11:51:32
 * golang中自定义的heap的使用: 小根堆
 */
package heap

type elem struct {
	val     uint64 // Value of this element.
	listIdx int    // Which list this element comes from.
}

type uint64Heap []elem

func (h uint64Heap) Len() int           { return len(h) }
func (h uint64Heap) Less(i, j int) bool { return h[i].val < h[j].val }

// 数组直接是引用，可直接修改
func (h uint64Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 指针接受者
func (h *uint64Heap) Push(x interface{}) {
	// 不是指针的话，添加不进去
	*h = append(*h, x.(elem))
}

// 指针接受者
func (h *uint64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
