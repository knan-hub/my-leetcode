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
 * @return {string}
 */
var tree2str = function (root) {
  // 方法一
  // if (!root) {
  //     return "";
  // }
  // let res = String(root.val);
  // if (root.left || root.right) {
  //     res += "(" + tree2str(root.left) + ")";
  // }
  // if (root.right) {
  //     res += "(" + tree2str(root.right) + ")";
  // }
  // return res;

  // 方法二
  if (!root) {
    return "";
  }
  let res = "";
  function search(node) {
    if (!node) {
      return;
    }
    res += String(node.val);
    if (node.left || node.right) {
      res += "(";
      search(node.left);
      res += ")";
    }
    if (node.right) {
      res += "(";
      search(node.right);
      res += ")";
    }
  }
  search(root);
  return res;
};
