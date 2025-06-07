/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function (nums1, nums2) {
  const memo = new Map();
  for (let i = 0; i < nums1.length; i++) {
    const num = nums1[i];
    memo.set(num, 1);
  }
  const resultMemo = new Map();
  for (let i = 0; i < nums2.length; i++) {
    const num = nums2[i];
    if (memo.has(num)) {
      resultMemo.set(num, 1);
    }
  }
  const result = [];
  for (const key of resultMemo.keys()) {
    result.push(key);
  }
  return result;
};
