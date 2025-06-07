/**
 * @param {number[]} arr
 * @param {number} k
 * @return {number}
 */
var findKthPositive = function (arr, k) {
  // 方法一
  // if (k <= arr[0] - 1) {
  //     return k;
  // }
  // k -= arr[0] - 1
  // for (let i = 1; i < arr.length; i++) {
  //     const missing = arr[i] - arr[i - 1] - 1;
  //     if (missing >= k) {
  //         return arr[i - 1] + k;
  //     }
  //     k -= missing;
  // }
  // return arr[arr.length - 1] + k;

  // 方法二
  let left = 0;
  let right = arr.length - 1;
  while (left <= right) {
    const middle = (left + right) >> 1;
    const missing = arr[middle] - (middle + 1);
    if (missing < k) {
      left = middle + 1;
    } else {
      right = middle - 1;
    }
  }
  return left + k;
};
