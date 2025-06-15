/**
 * @param {string} s
 * @return {string[]}
 */
var letterCasePermutation = function (s) {
  const result = [];
  function dfs(index, path) {
    if (index === s.length) {
      result.push(path);
      return;
    }
    const char = s[index];
    if (/[a-zA-Z]/.test(char)) {
      dfs(index + 1, path + char.toLowerCase());
      dfs(index + 1, path + char.toUpperCase());
    } else {
      dfs(index + 1, path + char);
    }
  }
  dfs(0, "");
  return result;
};
