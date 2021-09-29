# todo

## 题目

## 解题
    lru.go ok

## LRU代码实现，C++实现
    https://zhuanlan.zhihu.com/p/133480542

## 算法题就像搭乐高：手把手带你拆解 LFU 算法
    https://mp.weixin.qq.com/s/oXv03m1J8TwtHwMJEZ1ApQ

1. 主要的数据结构
    HashMap<Integer, Integer> keyToVal
    存储key到value的映射
    HashMap<Integer, Integer> keyToFreq;
    存储key的freq
    HashMap<Integer, LinkedHashSet<Integer>> freqToKeys;
    int minFreq = 0;
    存储freq到对应的key的映射，用户key的过期

    LinkedHashSet顾名思义，是链表和哈希集合的结合体。链表不能快速访问链表节点，但是插入元素具有时序；哈希集合中的元素无序，但是可以对元素进行快速的访问和删除
    感觉c++或者go只能使用rb hash来实现

## Golang Gcache中的LRU和LFU
    https://www.jianshu.com/p/b5d29dc16906

    freq的记录将hash和两个list组合起来; 看上去比上面的java方式更简洁；但实现的是遍历随机时间过期淘汰，并不是按访问频次来淘汰

## LFU自己的思路
    通过hash存储key val
    通过hash存储freq 到 key的映射，然后这个key通过list来存储；做到插入和删除O1操作
    存储一个最小的freq
