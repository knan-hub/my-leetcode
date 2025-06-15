/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function (head) {
  if (!head) {
    return null;
  }
  let newHead = head;
  while (newHead.next) {
    const next = newHead.next;
    if (newHead.val === next.val) {
      newHead.next = next.next;
      continue;
    }
    newHead = next;
  }
  return head;
};
