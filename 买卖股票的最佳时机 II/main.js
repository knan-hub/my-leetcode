/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  // 方法一：贪心算法
  // let profit = 0;
  // for (let i = 1; i < prices.length; i++) {
  //     if (prices[i] > prices[i - 1]) {
  //         profit += prices[i] - prices[i - 1];
  //     }
  // }
  // return profit;

  // 方法二：动态规划 二维数组
  // const dp = Array.from({ length: prices.length }, () => [0, 0]);
  // // 第 0 天没股票，利润是 0
  // dp[0][0] = 0;
  // // 第 0 天买了股票，利润是 -prices[0]
  // dp[0][1] = -prices[0];
  // for (let i = 1; i < prices.length; i++) {
  //     // 第 i 天没持股 = max(昨天也没持股，昨天持股今天卖出)
  //     dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
  //     // 第 i 天持股 = max(昨天也持股，昨天没持股今天买入)
  //     dp[i][1] = Math.max(dp[i - 1][1], dp[i - 1][0] - prices[i]);
  // }
  // return dp[prices.length - 1][0];

  // 方法三：动态规划 滚动变量
  let cash = 0; // 手里没股票
  let hold = -prices[0]; // 手里有股票
  for (let i = 1; i < prices.length; i++) {
    [cash, hold] = [
      Math.max(cash, hold + prices[i]),
      Math.max(hold, cash - prices[i]),
    ];
  }
  return cash;
};
