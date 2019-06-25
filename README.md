## 数据结构和算法的 Golang 实现

- [数组](arrays)
    - [支持动态扩容缩容的数组](arrays/dynamic.go)
    - [大小固定的有序数组](arrays/ordered.go)
    - [88. 合并两个有序数组](arrays/merge_test.go)
    - [15. 三数之和](arrays/threesum_test.go)
    - [169. 求众数(散列表法与排序法)](arrays/majority_test.go)
    - [41. 缺失的第一个正数](arrays/firstmiss_test.go)
    - [485. 最大连续1的个数](arrays/maxones_test.go)
    - [无序数组排序后的最大相邻差](arrays/maxdistance_test.go)
    - [238. 除自身以外数组的乘积](arrays/product_test.go)

- [链表](list)
    - [单链表](list/slist.go)
    - [循环链表](list/clist.go)
    - [双向链表](list/dlist.go)
    - [206. 反转链表](list/reversal_test.go)
    - [876. 链表的中间节点](list/mid_test.go)
    - [合并两个有序链表](list/merge2_test.go)
    - [141. 链表中是否有环](list/hasring_test.go)
    - [142. 环形链表 II](list/hasring2_test.go)
    - [23. 合并 K 个排序链表](list/mergek_test.go)
    - [链表的倒数第 K 个节点](list/lastk_test.go)
    - [链表中环的长度](list/ringlen_test.go)
    - [19. 删除链表的倒数第N个节点](list/removenth_test.go)

- [栈](stack)
    - [顺序栈](stack/arraystack.go)
    - [链式栈](stack/liststack.go)
    - [20. 有效的括号](stack/parentheses_test.go)
    - [最小栈](stack/minstack.go)

- [队列](queue)
    - [顺序队列](queue/arrayqueue.go)
    - [链式队列](queue/listqueue.go)
    - [循环队列](queue/cyclequeue.go)

- [递归](recursion)
    - [509. 斐波那契数](recursion/fib_test.go)
    - [求阶乘](recursion/factorial_test.go)
    - [求全排列](recursion/permutate_test.go)
    - [50. Pow(x, n)](recursion/pow_test.go)
    - [22. 括号生成](recursion/genparenthesis_test.go)
    - [最大公约数](recursion/divisor_test.go)

- [排序](sort)
    - [冒泡排序](sort/bubble.go)
    - [插入排序](sort/insertion.go)
    - [选择排序](sort/selection.go)
    - [归并排序](sort/merge.go)
    - [快速排序](sort/quick.go)
    - [O(n) 时间复杂度内找到一组数据的第 K 大元素](sort/kth_test.go)

- [二分查找](binarysearch)
    - [简单二分查找](binarysearch/bsearch_test.go)
    - [查找第一个值等于给定值的元素](binarysearch/bsearch1_test.go)
    - [查找最后一个值等于给定值的元素](binarysearch/bsearch2_test.go)
    - [查找第一个大于等于给定值的元素](binarysearch/bsearch3_test.go)
    - [查找最后一个小于等于给定值的元素](binarysearch/bsearch4_test.go)
    - [69. 求 x 的平方根](binarysearch/sqrt_test.go)

- [散列表](hashtable)
    - [基于链表法解决冲突的散列表](hashtable/listhashtable.go)
    - [有效的字母异位词](hashtable/isanagram_test.go)

- [字符串](string)
    - [字符串反转](string/reverse_test.go)

- [二叉树](binarytree)
    - [二叉搜索树](binarytree/bst.go)
    - [144. 前序遍历](binarytree/preorder_test.go)
    - [94. 中序遍历](binarytree/inorder_test.go)
    - [145. 后序遍历](binarytree/postorder_test.go)
    - [102. 层序遍历](binarytree/levelorder_test.go)
    - [104. 二叉树的最大深度](binarytree/maxdepth_test.go)
    - [111. 二叉树的最小深度](binarytree/mindepth_test.go)
    - [114. 二叉树展开为链表](binarytree/flatten_test.go)
    - [98. 验证二叉搜索树](binarytree/validbst_test.go)
    - [226. 翻转二叉树](binarytree/inverttree_test.go)
    - [235. 二叉搜索树的最近公共祖先](binarytree/bstlca_test.go)
    - [236. 二叉树的最近公共祖先](binarytree/lca_test.go)
    - [124. 二叉树中的最大路径和](binarytree/maxpathsum_test.go)

- [堆](heap)
    - [小顶堆](heap/minheap.go)
    - [大顶堆](heap/maxheap.go)
    - [优先级队列](heap/priorityqueue.go)
    - [堆排序](heap/sort.go)

- [贪心算法](greedy)
    - [122. 买卖股票的最佳时机 II](greedy/maxprofit_test.go)

- [回溯算法](backtrack)
    - [51. N皇后](backtrack/queens_test.go)
    - [52. N皇后 II](backtrack/totalq_test.go)
    - [36. 有效的数独](backtrack/validsudoku_test.go)
    - [37. 解数独](backtrack/sudoku_test.go)
    - [79. 单词搜索](backtrack/wsearch_test.go)

- [Trie](trie)
    - [208. 实现 Trie (前缀树)](trie/trie.go)
    - [212. 单词搜索 II](trie/findwords_test.go)

- [位运算](bit)
    - [191. 位1的个数](bit/count_test.go)
    - [231. 2的幂](bit/powerof2_test.go)
    - [338. 比特位计数](bit/countbits_test.go)
    - [52. N皇后 II（位运算）](bit/totalq_test.go)