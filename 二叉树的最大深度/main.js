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
var maxDepth = function (root) {
  // 方法一
  // if (!root) {
  //     return 0;
  // }
  // if (!root.left && !root.right) {
  //     return 1;
  // }
  // return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));

  // 方法二
  if (!root) {
    return 0;
  }
  const queue = [root];
  let depth = 0;
  while (queue.length) {
    const levelSize = queue.length;
    for (let i = 0; i < levelSize; i++) {
      const node = queue.shift();
      if (node.left) {
        queue.push(node.left);
      }
      if (node.right) {
        queue.push(node.right);
      }
    }
    depth++;
  }
  return depth;
};
