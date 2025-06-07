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
 * @return {boolean}
 */
var isBalanced = function (root) {
  // 方法一
  // if (!root) {
  //     return true;
  // }
  // const left = getHeight(root.left);
  // const right = getHeight(root.right);
  // if (Math.abs(left - right) > 1) {
  //     return false;
  // }
  // return isBalanced(root.left) && isBalanced(root.right);

  // 方法二
  return checkHeight(root) !== -1;
};

function getHeight(node) {
  if (!node) {
    return 0;
  }
  if (!node.left && !node.right) {
    return 1;
  }
  return 1 + Math.max(getHeight(node.left), getHeight(node.right));
}

function checkHeight(node) {
  if (!node) {
    return 0;
  }
  if (!node.left && !node.right) {
    return 1;
  }
  const left = checkHeight(node.left);
  if (left === -1) {
    return -1;
  }
  const right = checkHeight(node.right);
  if (right === -1) {
    return -1;
  }
  if (Math.abs(left - right) > 1) {
    return -1;
  }
  return 1 + Math.max(left, right);
}
