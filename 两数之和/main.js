/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  const mapping = new Map();
  for (let i = 0; i < nums.length; i++) {
    const num = target - nums[i];
    if (mapping.has(num)) {
      return [i, mapping.get(num)];
    }
    mapping.set(nums[i], i);
  }
  return [];
};
