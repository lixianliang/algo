# tree


## 串联每一层的兄弟节点 ok
1. 土办法 将每一层使用map存储起来
level_range.go 
1. 利用上一层已经ready来进行操作
   涉及到树和链表的指针操作

## 寻找最深一层所有根节点的数值之和 ok
    deep_level_sum.go 
    递归解决

## 剑指 Offer 36. 二叉搜索树与双向链表

## 二叉树的子结构

## NC15 求二叉树的层序遍历 ok
    使用go的sliec存储要遍历的节点，通过low high游标来控制循环

## NC13 二叉树的最大深度 ok
    使用递归调用

## NC45 实现二叉树先序，中序和后序遍历 ok
    三种递归遍历树的方式

## NC8 二叉树根节点到叶子节点和为指定值的路径
    递归遍历 
    通过slice传递和回溯
