/**
 * @param {number[]} places
 * @return {boolean}
 */
var checkDynasty = function (places) {
  // 方法一
  // places.sort((a, b) => a - b);
  // let zeroCount = 0;
  // let gaps = 0;
  // for (const place of places) {
  //     if (!place) {
  //         zeroCount++;
  //     }
  // }
  // for (let i = zeroCount; i < places.length - 1; i++) {
  //     if (places[i] === places[i + 1]) {
  //         return false;
  //     }
  //     gaps += places[i + 1] - places[i] - 1;
  // }
  // return gaps <= zeroCount;

  // 方法二
  const memo = {};
  let minNum = 14;
  let maxNum = 0;
  for (const place of places) {
    if (!place) {
      continue;
    }
    if (memo[place]) {
      return false;
    }
    minNum = Math.min(minNum, place);
    maxNum = Math.max(maxNum, place);
    memo[place] = 1;
  }
  return maxNum - minNum < 5;
};
