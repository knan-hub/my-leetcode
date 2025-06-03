/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function (root, p, q) {
  // 方法一
  // if (!root || root === p || root === q) {
  //     return root;
  // }
  // const left = lowestCommonAncestor(root.left, p, q);
  // const right = lowestCommonAncestor(root.right, p, q);
  // if (left && right) {
  //     return root;
  // }
  // if (left) {
  //     return left;
  // }
  // return right;

  // 方法二
  // const search = (node) => {
  //     if (!node || node === p || node === q) {
  //         return node;
  //     }
  //     const left = search(node.left);
  //     const right = search(node.right);
  //     if (left && right) {
  //         return node;
  //     }
  //     if (left) {
  //         return left;
  //     }
  //     return right;
  // };
  // return search(root);

  // 方法三
  const parentMap = new Map();
  const stack = [root];
  parentMap.set(root, null);
  while (stack.length && !(p && parentMap.get(p) && q && parentMap.get(q))) {
    const node = stack[stack.length - 1];
    stack.pop();
    if (node.left) {
      parentMap.set(node.left, node);
      stack.push(node.left);
    }
    if (node.right) {
      parentMap.set(node.right, node);
      stack.push(node.right);
    }
  }
  const ancestors = new Map();
  while (p) {
    ancestors.set(p, true);
    p = parentMap.get(p);
  }
  while (q) {
    if (ancestors.get(q)) {
      return q;
    }
    q = parentMap.get(q);
  }
  return null;
};
