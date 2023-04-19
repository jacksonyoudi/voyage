package skiplist

import (
	"math/rand"
)

type node struct {
	// 长度对应为当前节点的高度
	nexts []*node
	key   int
	val   int
}

type Skiplist struct {
	head *node
}

// 根据 key 读取 val，
//第二个 bool flag 反映 key 在 skiplist 中是否存在
func (s *Skiplist) Get(key int) (int, bool) {

	// 不符合 GO的开发风格
	if _node := s.search(key); _node != nil {
		return _node.val, true
	}
	return -1, false

}

// 从跳表中检索 key 对应的 node
func (s *Skiplist) search(key int) *node {
	// 每次检索都是从头部开始
	move := s.head
	// 每次检索从最大高度出发,直到来到首层
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 在每一层中持续向右遍历，
		//发现 next比 key小, 说明key在 下一个next后,
		//故 move = move.nexts[level]
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}

		// 如果下一个key值相等, 就返回 node
		if move.nexts[level] != nil && move.nexts[level].key == key {
			return move.nexts[level]
		}

		// 如果当前层没有找到目标, 则层数减1, 继续向下
	}
	return nil
}

//roll骰子 决定一个待插入的新节点在skiplist中最高层对应的index
func (s *Skiplist) roll() int {
	var level int
	// 每次投出1 则层数 +1
	for rand.Int() > 0 {
		level++
	}
	return level
}

func (s *Skiplist) Put(key, val int) {
	// 假如 kv对已存在，则直接对值进行更新并返回
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}

	// roll 出新节点的高度
	level := s.roll()

	// 新节点高度超出跳表最大高度, 则需要对齐高度进行补充
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}

	// 创建出新的节点
	newNode := &node{
		key:   key,
		val:   val,
		nexts: make([]*node, level),
	}

	// 从头节点的最高层出发
	move := s.head
	for l := level; l >= 0; l-- {
		// 向右遍历，直到右侧节点不存在或者 key 值大于 key
		for move.nexts[l] != nil && move.nexts[l].key < key {
			move = move.nexts[level]
		}
		// 调整指针关系, 完成新节点的插入
		newNode.nexts[l] = move.nexts[l]
		move.nexts[l] = newNode
	}

}

// 根据key从跳表中删除对应的节点
func (s *Skiplist) Del(key int) {
	// 如果 kv 对不存在，则无需删除直接返回
	if _node := s.search(key); _node == nil {
		return
	}

	// 从头节点的最高层出发
	move := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 向右遍历, 直到右侧节点不存在或者key值大于等于key
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}
		// 右侧节点不存在 或者key值大于target, 则直接跳过
		if move.nexts[level] == nil || move.nexts[level].key > key {
			continue
		}

		// 走到此处意味着右侧节点的key值必然等于key, 则调整指针引用
		move.nexts[level] = move.nexts[level].nexts[level]
	}

}
