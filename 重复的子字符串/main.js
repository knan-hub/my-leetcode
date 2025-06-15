/**
 * @param {string} s
 * @return {boolean}
 */
var repeatedSubstringPattern = function (s) {
  // 方法一
  // const ss = (s + s).slice(1, -1);
  // return ss.includes(s);

  // 方法二
  const n = s.length;
  for (let i = 1; i <= Math.floor(n / 2); i++) {
    if (n % i !== 0) {
      continue;
    }
    const pattern = s.slice(0, i);
    const repeated = pattern.repeat(n / i);
    if (repeated === s) {
      return true;
    }
  }
  return false;
};
