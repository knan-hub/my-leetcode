/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function (head) {
  if (!head || !head.next) {
    return true;
  }
  const dummy = new ListNode(-1, head);
  let slow = dummy;
  let fast = dummy;
  while (fast && fast.next) {
    slow = slow.next;
    fast = fast.next.next;
  }
  let secondHalf = slow.next;
  slow.next = null;
  let reversed = null;
  while (secondHalf) {
    const next = secondHalf.next;
    secondHalf.next = reversed;
    reversed = secondHalf;
    secondHalf = next;
  }
  while (head && reversed) {
    if (head.val !== reversed.val) {
      return false;
    }
    head = head.next;
    reversed = reversed.next;
  }
  return true;
};
