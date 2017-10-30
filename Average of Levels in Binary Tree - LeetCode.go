/*
Average of Levels in Binary Tree - LeetCodeGiven a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

Example 1:

Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].



Note:

The range of node's value is in the range of 32-bit signed integer.

*/
package averageOfLevels

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func dfs(root *TreeNode, buffer *[][]int, deep int){
  if root!=nil{
    if len(*buffer) <= deep{
      var tmp = []int{0, 0}
      *buffer = append(*buffer, tmp)
    }
    (*buffer)[deep][0] += root.Val
    (*buffer)[deep][1] += 1
    dfs(root.Left, buffer, deep+1)
    dfs(root.Right, buffer, deep+1)
  }
}
func averageOfLevels(root *TreeNode) []float64 {
  var buffer = [][]int{}
  var result = []float64{}
  dfs(root, &buffer, 0)
  for _, v:= range(buffer){
    result = append(result, float64(v[0])/float64(v[1]))
  }
  return result
}