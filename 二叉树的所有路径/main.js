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
 * @return {string[]}
 */
var binaryTreePaths = function (root) {
  // 方法一
  // if (!root) {
  //     return null;
  // }
  // const result = [];
  // const search = (node, path) => {
  //     if (!node.left && !node.right) {
  //         result.push(path);
  //         return;
  //     }
  //     if (node.left) {
  //         search(node.left, path + "->" + node.left.val);
  //     }
  //     if (node.right) {
  //         search(node.right, path + "->" + node.right.val);
  //     }
  // };
  // search(root, `${root.val}`);
  // return result;

  // 方法二
  if (!root) {
    return null;
  }
  const result = [];
  const nodeQueue = [root];
  const pathQueue = [`${root.val}`];
  for (let i = 0; i < nodeQueue.length; i++) {
    const node = nodeQueue[i];
    const path = pathQueue[i];
    if (!node.left && !node.right) {
      result.push(path);
      continue;
    }
    if (node.left) {
      nodeQueue.push(node.left);
      pathQueue.push(path + "->" + node.left.val);
    }
    if (node.right) {
      nodeQueue.push(node.right);
      pathQueue.push(path + "->" + node.right.val);
    }
  }
  return result;
};
