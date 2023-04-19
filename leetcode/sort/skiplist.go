package main

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 16  // 最大层数
	p        = 0.5 // 随机概率
)

type node struct {
	key     int
	value   interface{}
	forward []*node
	level   int
}

type skipList struct {
	head   *node
	length int
}

func newNode(key int, value interface{}, level int) *node {
	return &node{
		key:     key,
		value:   value,
		forward: make([]*node, level),
	}
}

func newSkipList() *skipList {
	return &skipList{
		head:   newNode(0, nil, maxLevel),
		length: 0,
	}
}

func (sl *skipList) randomLevel() int {
	level := 1
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

func (sl *skipList) insert(key int, value interface{}) {
	update := make([]*node, maxLevel)
	x := sl.head
	for i := maxLevel - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		x.value = value
	} else {
		level := sl.randomLevel()
		if level > sl.head.level {
			for i := sl.head.level; i < level; i++ {
				update[i] = sl.head
			}
			sl.head.level = level
		}
		x = newNode(key, value, level)
		for i := 0; i < level; i++ {
			x.forward[i] = update[i].forward[i]
			update[i].forward[i] = x
		}
		sl.length++
	}
}

func (sl *skipList) delete(key int) {
	update := make([]*node, maxLevel)
	x := sl.head
	for i := maxLevel - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		for i := 0; i < sl.head.level; i++ {
			if update[i].forward[i] != x {
				break
			}
			update[i].forward[i] = x.forward[i]
		}
		for sl.head.level > 1 && sl.head.forward[sl.head.level-1] == nil {
			sl.head.level--
		}
		sl.length--
	}
}

func (sl *skipList) search(key int) interface{} {
	x := sl.head
	for i := sl.head.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		return x.value
	}
	return nil
}

func main() {
	sl := newSkipList()
	sl.insert(1, "one")
	sl.insert(2, "two")
	sl.insert(3, "three")
	fmt.Println(sl)
}
