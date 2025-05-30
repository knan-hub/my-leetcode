/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  // 方法一
  // let maxProfit = 0;
  // for (let i = 0; i < prices.length; i++) {
  //     let profit = 0;
  //     for (let j = i + 1; j < prices.length; j++) {
  //         if (prices[j] > prices[i]) {
  //             profit = Math.max(profit, prices[j] - prices[i]);
  //         }
  //     }
  //     maxProfit = Math.max(maxProfit, profit);
  // }
  // return maxProfit;

  // 方法二
  // let minPrice = Infinity;
  // let maxPrice = 0;
  // for (const price of prices) {
  //     if (price < minPrice) {
  //         minPrice = price;
  //     } else {
  //         maxPrice = Math.max(maxPrice, price - minPrice);
  //     }
  // }
  // return maxPrice;

  // 方法三
  // const dp = Array.from({ length: prices.length }, () => [0, 0]);
  // dp[0][0] = 0;
  // dp[0][1] = -prices[0];
  // for (let i = 1; i < prices.length; i++) {
  //     dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
  //     dp[i][1] = Math.max(dp[i - 1][1], -prices[i]);
  // }
  // return dp[prices.length - 1][0];

  // 方法四
  let cash = 0;
  let hold = -prices[0];
  for (let i = 1; i < prices.length; i++) {
    cash = Math.max(cash, hold + prices[i]);
    hold = Math.max(hold, -prices[i]);
  }
  return cash;
};
