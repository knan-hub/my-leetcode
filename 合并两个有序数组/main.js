/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  let indexM = m - 1;
  let indexN = n - 1;
  let index = m + n - 1;
  while (indexM >= 0 && indexN >= 0) {
    if (nums1[indexM] < nums2[indexN]) {
      nums1[index] = nums2[indexN];
      indexN--;
    } else {
      nums1[index] = nums1[indexM];
      indexM--;
    }
    index--;
  }
  while (indexN >= 0) {
    nums1[index] = nums2[indexN];
    index--;
    indexN--;
  }
};
