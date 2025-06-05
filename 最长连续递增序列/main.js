/**
 * @param {number[]} nums
 * @return {number}
 */
var findLengthOfLCIS = function (nums) {
  if (!nums.length) {
    return 0;
  }
  let i = 0;
  let maxSize = 1;
  for (let j = 1; j < nums.length; j++) {
    if (nums[j] > nums[j - 1]) {
      maxSize = Math.max(maxSize, j - i + 1);
    } else {
      i = j;
    }
  }
  return maxSize;
};
