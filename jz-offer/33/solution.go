package main

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func verifyPostorder(postorder []int) bool {
	return dfs(postorder, 0, len(postorder)-1)
}

func dfs(nums []int, l, r int) bool {
	if l >= r {
		return true
	}
	last := nums[r]
	k := l //右子树的开始下标
	for k < r && nums[k] < last {
		k++
	}
	for j := k; j < r; j++ {
		if nums[j] < last {
			return false
		}
	}
	return dfs(nums, l, k-1) && dfs(nums, k, r-1)
}

func main() {
	//a := []int{4, 8, 6, 12, 16, 14, 10}
	a := []int{5, 4, 3, 2, 1}
	//a := []int{1, 3, 2, 6, 5}
	println(verifyPostorder(a))
}
