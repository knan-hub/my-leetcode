/**
 * @param {number[]} sales
 * @return {number}
 */
var maxSales = function (sales) {
  // 方法一
  // if (!sales.length) {
  //     return 0;
  // }
  // let maxSale = sales[0];
  // let currentSale = sales[0];
  // for (let i = 1; i < sales.length; i++) {
  //     if (currentSale < 0) {
  //         currentSale = sales[i];
  //     } else {
  //         currentSale += sales[i];
  //     }
  //     if (currentSale > maxSale) {
  //         maxSale = currentSale;
  //     }
  // }
  // return maxSale;

  // 方法二
  let max = sales[0];
  for (let i = 1; i < sales.length; i++) {
    if (sales[i] + sales[i - 1] > sales[i]) {
      sales[i] += sales[i - 1];
    }
    if (sales[i] > max) {
      max = sales[i];
    }
  }
  return max;
};
