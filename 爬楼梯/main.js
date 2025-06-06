/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  // 方法一
  // if (n === 1) {
  //     return 1;
  // }
  // if (n === 2) {
  //     return 2;
  // }
  // let i = 1;
  // let j = 2;
  // for (let k = 3; k <= n; k++) {
  //     const num = i + j;
  //     i = j;
  //     j = num;
  // }
  // return j;

  // 方法二
  // if (n === 1) {
  //     return 1;
  // }
  // if (n === 2) {
  //     return 2;
  // }
  // return climbStairs(n - 1) + climbStairs(n - 2);

  // 方法三
  const memo = new Map();
  return helper(n, memo);
};

function helper(n, memo) {
  if (n === 1) {
    return 1;
  }
  if (n === 2) {
    return 2;
  }
  if (memo.get(n)) {
    return memo.get(n);
  }
  memo.set(n, helper(n - 1, memo) + helper(n - 2, memo));
  return memo.get(n);
}
