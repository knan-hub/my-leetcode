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
var isSymmetric = function (root) {
  // 方法一
  // const check = (left, right) => {
  //     if (!left && !right) {
  //         return true;
  //     }
  //     if (!left || !right) {
  //         return false;
  //     }
  //     if (left.val !== right.val) {
  //         return false;
  //     }
  //     return check(left.left, right.right) && check(left.right, right.left);
  // };
  // return check(root.left, root.right);

  // 方法二
  const queue = [root.left, root.right];
  while (queue.length) {
    const left = queue.shift();
    const right = queue.shift();
    if (!left && !right) {
      continue;
    }
    if (!left || !right) {
      return false;
    }
    if (left.val !== right.val) {
      return false;
    }
    queue.push(...[left.left, right.right, left.right, right.left]);
  }
  return true;
};
