/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function (num1, num2) {
  if (num1.length < num2.length) {
    [num1, num2] = [num2, num1];
  }
  num2 = "0".repeat(num1.length - num2.length) + num2;
  let carry = 0;
  let result = [];
  for (let i = num1.length - 1; i >= 0; i--) {
    const a = parseInt(num1[i]);
    const b = parseInt(num2[i]);
    const sum = a + b + carry;
    carry = Math.floor(sum / 10);
    result.push(sum % 10);
  }
  if (carry > 0) {
    result.push(carry);
  }
  return result.reverse().join("");
};
