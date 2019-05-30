## 数据结构和算法的 Golang 实现

- [数组](arrays)
    - [支持动态扩容缩容的数组](arrays/dynamic.go)
    - [大小固定的有序数组](arrays/ordered.go)
    - [合并两个有序数组](arrays/merge_test.go)
    - [三数之和](arrays/threesum_test.go)
    - [求众数(散列表法与排序法)](arrays/majority_test.go)
    - [缺失的第一个正数](arrays/firstmiss_test.go)

- [链表](list)
    - [单链表](list/slist.go)
    - [循环链表](list/clist.go)
    - [双向链表](list/dlist.go)
    - [单链表反转](list/reversal_test.go)
    - [求链表的中间结点](list/mid_test.go)
    - [合并两个有序链表](list/merge2_test.go)
    - [链表中是否有环](list/hasring_test.go)
    - [合并 K 个排序链表](list/mergek_test.go)

- [栈](stack)
    - [顺序栈](stack/arraystack.go)
    - [链式栈](stack/liststack.go)
    - [有效的括号](stack/parentheses_test.go)

- [队列](queue)
    - [顺序队列](queue/arrayqueue.go)
    - [链式队列](queue/listqueue.go)
    - [循环队列](queue/cyclequeue.go)

- [递归](recursion)
    - [斐波那契数列求值](recursion/fib_test.go)
    - [求阶乘](recursion/factorial_test.go)
    - [求全排列](recursion/permutate_test.go)

- [排序](sort)
    - [冒泡排序](sort/bubble.go)
    - [插入排序](sort/insertion.go)
    - [选择排序](sort/selection.go)
    - [归并排序](sort/merge.go)
    - [快速排序](sort/quick.go)
    - [O(n) 时间复杂度内找到一组数据的第 K 大元素](sort/kth_test.go)