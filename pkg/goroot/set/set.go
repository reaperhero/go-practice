package set

import "sync"

type Set struct {
	m            map[int]struct{} // 用字典来实现，因为字段键不能重复
	len          int          // 集合的大小
	sync.RWMutex              // 锁，实现并发安全
}

// 新建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
	s.len = len(s.m)       // 重新计算元素数量
}

// 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()

	// 集合没元素直接返回
	if s.len == 0 {
		return
	}

	delete(s.m, item) // 实际从字典删除这个键
	s.len = len(s.m)  // 重新计算元素数量
}

// 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// 集合是够为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{} // 字典重新赋值
	s.len = 0                // 大小归零
}


// 将集合转化为切片
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}