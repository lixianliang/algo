# 递归

## 数据结构：图的存储结构之邻接矩阵
    图的存储
    定点数组，可以记录对应的入度和出度；
    邻接矩阵：使用NxN二维数组可以记录次有向图，对角线都是0,
    入度：指向次节点的变；从此节点出去的边；行的总数为入度，列的总数为出度

## 数字全排列
    有重复的情况需要做排序



## NC42 有重复项数字的所有排列
    listall.go  ac

### 方法
    利用递归模拟概率选取

## NC43 没有重复项数字的所有排列
    listall1.go  ac

## N皇后问题
    https://leetcode-cn.com/problems/eight-queens-lcci/
1. 使用二维数组模拟棋盘，0表示空白，1表示摆放棋盘
    queue.go 自测ok

## 数组的全排列
    https://cloud.tencent.com/developer/article/1176997

## LeetCode 31：递归、回溯、八皇后、全排列一篇文章全讲清楚
    https://cloud.tencent.com/developer/article/1594817

    递归思想

## LeetCode52题，别再问我N皇后问题了
    https://www.cnblogs.com/techflow/p/12782020.html

## 经典算法之八皇后问题
    https://zhuanlan.zhihu.com/p/99209213

## 什么是拓扑排序（Topological Sorting）
    https://www.jianshu.com/p/b59db381561a


## LeetCode207. Course Schedule课程表
    https://blog.csdn.net/LTuantuan/article/details/102654107?utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control
1. 思路
    1. 计算入度表
    1. 先将入度为0的节点放到que中
    1. 然后逐个出队列，并减少对应的出度（对应的节点需要入度减少），同时将入度为0的节点放到queue；对应的节点数减1
    1. 当队列没有节点，且没有入度为1的节点，则是有向无环图
    

## Golang Leetcode 207. Course Schedule.go
    https://blog.csdn.net/anakinsun/article/details/89013244
    leetcode-golang 里面包含leetcode专栏代码


