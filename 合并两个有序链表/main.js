/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function (list1, list2) {
  // 方法一
  // if (list1 == null) return list2;
  // if (list2 == null) return list1;

  // if (list1.val <= list2.val) {
  //     list1.next = mergeTwoLists(list1.next, list2);
  //     return list1;
  // }

  // list2.next = mergeTwoLists(list1, list2.next);
  // return list2;

  // 方法二
  const head = new ListNode(-1);
  let cur = head;

  while (list1 != null && list2 != null) {
    if (list1.val <= list2.val) {
      cur.next = list1;
      list1 = list1.next;
    } else {
      cur.next = list2;
      list2 = list2.next;
    }
    cur = cur.next;
  }

  if (list1 == null) {
    cur.next = list2;
  } else {
    cur.next = list1;
  }

  return head.next;
};
