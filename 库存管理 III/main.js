/**
 * @param {number[]} stock
 * @param {number} cnt
 * @return {number[]}
 */
var inventoryManagement = function (stock, cnt) {
  sort(stock, 0, stock.length - 1, cnt);
  return stock.slice(0, cnt);
};

function sort(stock, left, right, cnt) {
  if (left >= right) {
    return;
  }
  const index = quickSort(stock, left, right);
  if (index === cnt) {
    return;
  } else if (index > cnt) {
    sort(stock, left, index - 1, cnt);
  } else {
    sort(stock, index + 1, right, cnt);
  }
}

function quickSort(stock, left, right) {
  let newLeft = left;
  let newRight = right;
  const temp = stock[newLeft];
  while (newLeft < newRight) {
    while (newLeft < newRight && stock[newRight] >= temp) {
      newRight--;
    }
    stock[newLeft] = stock[newRight];
    while (newLeft < newRight && stock[newLeft] <= temp) {
      newLeft++;
    }
    stock[newRight] = stock[newLeft];
  }
  stock[newLeft] = temp;
  return newLeft;
}
