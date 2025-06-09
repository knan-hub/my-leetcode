/**
 * @param {string} s
 * @return {number}
 */
var longestPalindrome = function (s) {
  const memo = {};
  for (const char of s) {
    if (memo[char]) {
      memo[char]++;
    } else {
      memo[char] = 1;
    }
  }
  let length = 0;
  let hasOdd = false;
  for (const key of Object.keys(memo)) {
    if (memo[key] % 2 === 0) {
      length += memo[key];
    } else {
      length += memo[key] - 1;
      hasOdd = true;
    }
  }
  if (hasOdd) {
    length++;
  }
  return length;
};
