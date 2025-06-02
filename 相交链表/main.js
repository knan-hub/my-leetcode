/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function (headA, headB) {
  let newHeadA = headA;
  let newHeadB = headB;
  while (newHeadA !== newHeadB) {
    newHeadA = newHeadA ? newHeadA.next : headB;
    newHeadB = newHeadB ? newHeadB.next : headA;
  }
  return newHeadA;
};
