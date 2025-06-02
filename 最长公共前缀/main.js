/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  if (!strs.length) {
    return "";
  }
  const result = [];
  const str = strs[0];
  for (let i = 0; i < str.length; i++) {
    let isMatch = true;
    for (let j = 1; j < strs.length; j++) {
      const curStr = strs[j];
      if (curStr.length <= i || curStr[i] !== str[i]) {
        isMatch = false;
        break;
      }
    }
    if (!isMatch) {
      return result.join("");
    }
    result.push(str[i]);
  }
  return result.join("");
};
