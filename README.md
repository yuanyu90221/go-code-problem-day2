# go-code-problem-day2

## 題目描述

給定一個 linked list 的 head node，將指定區間內的節點反轉，並 return 新 linked list 的 head node。

指定區間之定義為: 第 m 個 node 到第 n 個 node 所形成的 linked list
Time complexity 要求: One-pass O(N)
Constraints:

m and n are positive integers
1 <= m < n <= length of the linked list
Example 1:

Input: Linked List = 6 -> 5 -> 4 -> 3 -> 2 -> 1, m = 2, n = 4
Output: 6 -> 3 -> 4 -> 5 -> 2 -> 1
Explanation: 將第二個 node 到第四個 node 所形成的 Linked List 反轉
Example 2:

Input: Linked List = 6 -> 5 -> 4 -> 3 -> 2 -> 1, m = 2, n = 6
Output: 6 -> 1 -> 2 -> 3 -> 4 -> 5
Explanation: 將第二個 node 到第六個 node 所形成的 Linked List 反轉

## 題目分析

給定一個 linked List 的 head node h, 將指定區間內的node反轉, 並且 return 新的 linked list的 head node

首先需要找出起始跟結束結點的指標 以及前一個結點

建立一個 stack 利用 stack FILO的特性來 重建 sub Linked lisk的順序

然後判斷是否 m=1 如果是 則 回傳sub LinkedList 否則回傳原本的 head



