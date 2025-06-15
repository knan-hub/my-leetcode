/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  // 方法一
  // const memo = {};
  // for (const num of nums) {
  //     if (memo[num]) {
  //         memo[num]++;
  //     } else {
  //         memo[num] = 1;
  //     }
  //     if (memo[num] > nums.length / 2) {
  //         return num;
  //     }
  // }
  // return -1;

  // 方法二
  let candidate = nums[0];
  let count = 1;
  for (let i = 1; i < nums.length; i++) {
    if (nums[i] === candidate) {
      count++;
    } else {
      count--;
      if (!count) {
        candidate = nums[i];
        count = 1;
      }
    }
  }
  return candidate;
};
