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
 * @return {number[]}
 */
var preorderTraversal = function (root) {
  // 方法一
  // const result = [];
  // const search = (node) => {
  //     if (!node) {
  //         return;
  //     }
  //     result.push(node.val);
  //     search(node.left);
  //     search(node.right);
  // };
  // search(root);
  // return result;

  // 方法二
  const result = [];
  const stack = [];
  let node = root;
  while (node || stack.length) {
    while (node) {
      result.push(node.val);
      stack.push(node);
      node = node.left;
    }
    node = stack.pop().right;
  }
  return result;
};
