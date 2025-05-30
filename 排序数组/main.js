/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArray = function (nums) {
  // 冒泡排序
  // for (let i = 0; i < nums.length; i++) {
  //     for (let j = i + 1; j < nums.length; j++) {
  //         if (nums[i] > nums[j]) {
  //             [nums[i], nums[j]] = [nums[j], nums[i]];
  //         }
  //     }
  // }
  // return nums;

  // 插入排序
  // for (let i = 1; i < nums.length; i++) {
  //     let j = i - 1;
  //     const num = nums[i];
  //     while (j >= 0 && nums[j] > num) {
  //         nums[j + 1] = nums[j];
  //         j--;
  //     }
  //     nums[j + 1] = num;
  // }
  // return nums;

  // 选择排序
  // for (let i = 0; i < nums.length; i++) {
  //     let minIndex = i;
  //     for (let j = i + 1; j < nums.length; j++) {
  //         if (nums[j] < nums[minIndex]) {
  //             minIndex = j;
  //         }
  //     }
  //     if (minIndex != i) {
  //         [nums[i], nums[minIndex]] = [nums[minIndex], nums[i]];
  //     }
  // }
  // return nums;

  // 快速排序
  // const sort = (nums, left, right) => {
  //     if (left >= right) {
  //         return;
  //     }
  //     let newLeft = left;
  //     let newRight = right;
  //     let curNum = nums[newLeft];
  //     while (newLeft < newRight) {
  //         while (newLeft < newRight && nums[newRight] >= curNum) {
  //             newRight--;
  //         }
  //         nums[newLeft] = nums[newRight];
  //         while (newLeft < newRight && nums[newLeft] <= curNum) {
  //             newLeft++;
  //         }
  //         nums[newRight] = nums[newLeft];
  //     }
  //     nums[newLeft] = curNum;
  //     sort(nums, left, newLeft - 1);
  //     sort(nums, newLeft + 1, right);
  // };
  // sort(nums, 0, nums.length - 1);
  // return nums;

  // 归并排序
  const sort = (nums, left, right) => {
    if (left >= right) {
      return;
    }
    const middle = (left + right) >> 1;
    sort(nums, left, middle);
    sort(nums, middle + 1, right);
    let newLeft = left;
    let newRight = middle + 1;
    let index = left;
    const tempNums = [];
    while (newLeft <= middle && newRight <= right) {
      if (nums[newLeft] <= nums[newRight]) {
        tempNums[index] = nums[newLeft];
        index++;
        newLeft++;
      } else {
        tempNums[index] = nums[newRight];
        index++;
        newRight++;
      }
    }
    while (newLeft <= middle) {
      tempNums[index] = nums[newLeft];
      index++;
      newLeft++;
    }
    while (newRight <= right) {
      tempNums[index] = nums[newRight];
      index++;
      newRight++;
    }
    for (let i = left; i <= right; i++) {
      nums[i] = tempNums[i];
    }
  };
  sort(nums, 0, nums.length - 1);
  return nums;
};
