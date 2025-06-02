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
var inorderTraversal = function (root) {
  // 方法一
  // const result = [];
  // const func = (node) => {
  //     if (!node) {
  //         return;
  //     }
  //     func(node.left);
  //     result.push(node.val);
  //     func(node.right);
  // };
  // func(root);
  // return result;

  // 方法二
  // const result = [];
  // const stack = [];
  // let cur = root;
  // while (cur || stack.length) {
  //     while (cur) {
  //         stack.push(cur);
  //         cur = cur.left;
  //     }
  //     cur = stack.pop();
  //     result.push(cur.val)
  //     cur = cur.right;
  // }
  // return result;

  // 方法三
  if (!root) {
    return [];
  }
  const result = [];
  result.push(...inorderTraversal(root.left));
  result.push(root.val);
  result.push(...inorderTraversal(root.right));
  return result;
};
