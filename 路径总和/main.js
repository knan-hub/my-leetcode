/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} targetSum
 * @return {boolean}
 */
var hasPathSum = function (root, targetSum) {
  // 方法一
  // if (!root) {
  //     return false;
  // }
  // function search(node, sum) {
  //     if (!node) {
  //         return false;
  //     }
  //     sum -= node.val;
  //     if (!node.left && !node.right) {
  //         return sum === 0;
  //     }
  //     return search(node.left, sum) || search(node.right, sum);
  // }
  // return search(root, targetSum);

  // 方法二
  if (!root) {
    return false;
  }
  if (!root.left && !root.right) {
    return targetSum === root.val;
  }
  return (
    hasPathSum(root.left, targetSum - root.val) ||
    hasPathSum(root.right, targetSum - root.val)
  );
};
