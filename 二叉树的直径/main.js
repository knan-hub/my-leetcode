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
 * @return {number}
 */
var diameterOfBinaryTree = function (root) {
  let maxDiameter = 0;
  const maxHeight = (node) => {
    if (!node) {
      return 0;
    }
    if (!node.left && !node.right) {
      return 1;
    }
    const left = maxHeight(node.left);
    const right = maxHeight(node.right);
    maxDiameter = Math.max(maxDiameter, left + right);
    return 1 + Math.max(left, right);
  };
  maxHeight(root);
  return maxDiameter;
};
