/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
var reverseBetween = function (head, left, right) {
  if (left === right) {
    return head;
  }
  let dummy = new ListNode(-1, head);
  let dummyHead = dummy;
  let count = left - 1;
  while (dummy.next && count > 0) {
    dummy = dummy.next;
    count--;
  }
  if (count > 0 || !dummy.next) {
    return head;
  }
  let newHead = dummy.next;
  let secondHead = newHead;
  let tail = null;
  count = right - left + 1;
  while (secondHead && count > 0) {
    const next = secondHead.next;
    secondHead.next = tail;
    tail = secondHead;
    secondHead = next;
    count--;
  }
  dummy.next = tail;
  newHead.next = secondHead;
  return dummyHead.next;
};
