/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} cnt
 * @return {ListNode}
 */
var trainingPlan = function (head, cnt) {
  if (!head) {
    return null;
  }
  let slow = new ListNode(-1, head);
  let fast = new ListNode(-1, head);
  while (cnt && fast) {
    fast = fast.next;
    cnt--;
  }
  while (fast) {
    fast = fast.next;
    slow = slow.next;
  }
  return slow;
};
