/*
Two Sum IV - Input is a BST - LeetCodeGiven a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

Input: 
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True




Example 2:

Input: 
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False



*/
package two_sum_iv_input_is_a_bst

func findTarget(root *TreeNode, k int) bool{
    return traverse(root, root, k)
 }
 
 func traverse(root *TreeNode, cur_node *TreeNode, k int) bool{
     var target = k - cur_node.Val
     if checknums(root, cur_node.Val, target){
         return true
     }else{
         if cur_node.Left != nil{
             if traverse(root, cur_node.Left, k){
                 return true
             }
         }
         if  cur_node.Right != nil{
             if traverse(root, cur_node.Right, k){
                 return true
             }
         }
     }
     return false
 }
 
 func checknums(root *TreeNode, src int, target int) bool{
     if root.Val == target && root.Val != src{
         return true
     }else if root.Val < target{
         if root.Right != nil{
             return checknums(root.Right, src, target)
         }
     }else if root.Val > target{
         if root.Left != nil{
             return checknums(root.Left, src, target)
         }
     }
     return false
 }