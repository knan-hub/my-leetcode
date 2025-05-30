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
var reverseList = function (head) {
  // 方法一
  // const newHead = new ListNode(-1);

  // while (head) {
  //     const next = head.next;
  //     head.next = newHead.next;
  //     newHead.next = head;
  //     head = next;
  // }

  // return newHead.next;

  // 方法二
  // let pre = null;
  // let cur = head;

  // while (cur) {
  //     const next = cur.next;
  //     cur.next = pre;
  //     pre = cur;
  //     cur = next;
  // }

  // return pre;

  // 方法三
  if (head == null || head.next == null) {
    return head;
  }

  const newHead = reverseList(head.next);
  head.next.next = head;
  head.next = null;

  return newHead;
};
