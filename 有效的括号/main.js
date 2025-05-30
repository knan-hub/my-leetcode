/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  // 方法一
  // const temp = [];
  // const map = {
  //     ')': '(',
  //     '}': '{',
  //     ']': '['
  // };

  // for (const char of s) {
  //     if (['(', '{', '['].includes(char)) {
  //         temp.push(char);
  //     } else if (temp.length == 0 || temp.pop() != map[char]) {
  //         return false;
  //     }
  // }

  // return temp.length == 0;

  // 方法二
  let preLen;
  do {
    preLen = s.length;
    s = s.replace("()", "").replace("[]", "").replace("{}", "");
  } while (s.length < preLen);
  return s.length == 0;
};
