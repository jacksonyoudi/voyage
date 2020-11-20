package Mapper

import "sync"

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

func (m *RWMap) Get(key int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, exist := m.m[key]
	return v, exist
}

func (m *RWMap) Set(key int, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

func (m *RWMap) Delete(key int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, key)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RLock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
