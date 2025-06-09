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
 * @param {number} k
 * @return {boolean}
 */
var findTarget = function (root, k) {
  // 方法一
  // const memo = {};
  // let isFind = false;
  // const search = (node) => {
  //     if (!node || isFind) {
  //         return;
  //     }
  //     const num = k - node.val;
  //     if (memo[num]) {
  //         isFind = true;
  //         return;
  //     }
  //     memo[node.val] = 1;
  //     search(node.left);
  //     search(node.right);
  // };
  // search(root);
  // return isFind;

  // 方法二
  // if (!root) {
  //     return false;
  // }
  // const memo = {};
  // const queue = [root];
  // while (queue.length) {
  //     const node = queue.pop();
  //     const num = k - node.val;
  //     if (memo[num]) {
  //         return true;
  //     }
  //     memo[node.val] = 1;
  //     if (node.left) {
  //         queue.push(node.left);
  //     }
  //     if (node.right) {
  //         queue.push(node.right);
  //     }
  // }
  // return false;

  // 方法三
  if (!root) {
    return false;
  }
  const leftStack = [];
  const rightStack = [];
  let node = root;
  while (node) {
    leftStack.push(node);
    node = node.left;
  }
  node = root;
  while (node) {
    rightStack.push(node);
    node = node.right;
  }
  const next = (stack, isLeft) => {
    if (!stack.length) {
      return null;
    }
    const node = stack.pop();
    let currentNode = node;
    if (isLeft) {
      currentNode = currentNode.right;
      while (currentNode) {
        leftStack.push(currentNode);
        currentNode = currentNode.left;
      }
    } else {
      currentNode = currentNode.left;
      while (currentNode) {
        rightStack.push(currentNode);
        currentNode = currentNode.right;
      }
    }
    return node;
  };
  let leftNode = next(leftStack, true);
  let rightNode = next(rightStack, false);
  while (leftNode && rightNode && leftNode !== rightNode) {
    const sum = leftNode.val + rightNode.val;
    if (sum === k) {
      return true;
    } else if (sum > k) {
      rightNode = next(rightStack, false);
    } else {
      leftNode = next(leftStack, true);
    }
  }
  return false;
};
