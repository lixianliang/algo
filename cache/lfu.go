package main

import (
	"container/list"
	"errors"
)

type Lfu struct {
	Len       int
	MinFreq   int
	Kvs       map[string]string
	KeyToFreq map[string]*FreqItem
	FreqToKey map[int]*list.List
}

type FreqItem struct {
	Freq int
	List *list.List
	Ele  *list.Element // freq key list
}

func NewLfu(max int) *Lfu {
	return &Lfu{
		Len:       max,
		MinFreq:   0,
		Kvs:       make(map[string]string, max),
		KeyToFreq: make(map[string]*FreqItem, max),
		FreqToKey: make(map[int]*list.List, max),
	}
}

func (lfu *Lfu) Get(key string) (error, string) {
	v, ok := lfu.Kvs[key]
	if !ok {
		return errors.New("not found"), ""
	}

	var freq *FreqItem
	freq, _ = lfu.KeyToFreq[key] // 需要key到 freq的映射
	/*if !ok {
		freq = FreqItem{Freq: 1, E: nil}
	} else {
	}*/
	freq.List.Remove(freq.Ele)
	if freq.List.Len() == 0 && lfu.MinFreq == freq.Freq {
		lfu.MinFreq += 1
	}
	freq.Freq += 1

	//key=freq
	l, ok := lfu.FreqToKey[freq.Freq] // 需要freq到key的映射
	if !ok {
		l = list.New()
		lfu.FreqToKey[freq.Freq] = l
	}

	freq.List = l
	freq.Ele = l.PushBack(key)

	return nil, v
}

func (lfu *Lfu) Set(key, val string) {
	v, ok := lfu.Kvs[key]
	if !ok { // 不存在
		if len(lfu.Kvs) >= lfu.Len {
			// 删除数据 通过最小MinFreq 找到对应要删除的key(freq到key的映射) 删除最老的minfreq节点 删除lfu.Kvs lfu.KeyToFreq
			l, ok := lfu.FreqToKey[lfu.MinFreq]
			if ok {
				e := l.Front()
				k := e.Value.(string)
				l.Remove(e)

				delete(lfu.Kvs, k)
				delete(lfu.KeyToFreq, k)
				lfu.Len--
			}
		}

		freq := &FreqItem{Freq: 1, List: nil, Ele: nil}
		l, ok := lfu.FreqToKey[freq.Freq]
		if !ok {
			l = list.New()
			lfu.FreqToKey[freq.Freq] = l
		}
		freq.Ele = l.PushBack(key)
		freq.List = l
		lfu.KeyToFreq[key] = freq

		lfu.Len++
		lfu.MinFreq = 1
		lfu.Kvs[key] = val
	} else {
		// 存在 更新
		if v != val {
			lfu.Kvs[key] = val
		}

		var freq *FreqItem
		freq, _ = lfu.KeyToFreq[key] // 需要key到 freq的映射
		/*if !ok {
			freq = FreqItem{Freq: 1, E: nil}
		} else {
		}*/
		freq.List.Remove(freq.Ele)
		if freq.List.Len() == 0 && lfu.MinFreq == freq.Freq {
			lfu.MinFreq += 1
		}
		freq.Freq += 1

		//key=freq
		l, ok := lfu.FreqToKey[freq.Freq] // 需要freq到key的映射
		if !ok {
			l = list.New()
			lfu.FreqToKey[freq.Freq] = l
		}

		freq.List = l
		freq.Ele = l.PushBack(key)
	}
}

func main() {
}
