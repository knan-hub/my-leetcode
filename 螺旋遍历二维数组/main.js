/**
 * @param {number[][]} array
 * @return {number[]}
 */
var spiralArray = function (array) {
  if (!array.length || !array[0].length) {
    return [];
  }
  const res = [];
  let top = 0;
  let bottom = array.length - 1;
  let left = 0;
  let right = array[0].length - 1;
  while (top <= bottom && left <= right) {
    for (let i = left; i <= right; i++) {
      res.push(array[top][i]);
    }
    top++;
    for (let i = top; i <= bottom; i++) {
      res.push(array[i][right]);
    }
    right--;
    if (top <= bottom) {
      for (let i = right; i >= left; i--) {
        res.push(array[bottom][i]);
      }
      bottom--;
    }
    if (left <= right) {
      for (let i = bottom; i >= top; i--) {
        res.push(array[i][left]);
      }
      left++;
    }
  }
  return res;
};
