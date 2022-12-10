package collections

// Set 用一个 map 来保存唯一的 key
type Set[T comparable] map[T]bool

// NewSet 根据传入的 size 构建一个空的 Set
func NewSet[T comparable](size int) Set[T] {
	return make(Set[T], size)
}

// Add 添加一个 key 到 set 中
func (s Set[T]) Add(key T) {
	s[key] = true
}

// Remove 在 set 中移除一个 key
func (s Set[T]) Remove(key T) {
	delete(s, key)
}

// Contains 检查某个 key 是否在 set 中
func (s Set[T]) Contains(key T) bool {
	return s[key]
}

// Difference 比较两个 set 的不同点
func (a Set[T]) Difference(b Set[T]) Set[T] {
	resultSet := NewSet[T](0)
	for key := range a {
		if !b.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// Union A union B
func (a Set[T]) Union(b Set[T]) Set[T] {
	small, large := smallLarge(a, b)

	for key := range small {
		large.Add(key)
	}
	return large
}

// Intersection A intersect B
func (a Set[T]) Intersection(b Set[T]) Set[T] {
	small, large := smallLarge(a, b)

	resultSet := NewSet[T](0)
	for key := range small {
		if large.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// smallLarge 根据其长度返回小集和大集
func smallLarge[T comparable](a, b Set[T]) (Set[T], Set[T]) {
	small, large := b, a
	if len(b) > len(a) {
		small, large = a, b
	}

	return small, large
}

// Equals 比较两个 set 是否相等
func (a Set[T]) Equals(b Set[T]) bool {
	return len(a.Difference(b)) == 0 && len(b.Difference(a)) == 0
}

// Equals 比较两个 set 是否相等
func (a Set[T]) ToSlice() []T {
	s := make([]T, 0)
	for key := range a {
		s = append(s, key)
	}
	return s
}

// SliceToSet
func SliceToSet[T comparable](s []T) Set[T] {
	set := NewSet[T](len(s))
	for _, item := range s {
		set.Add(item)
	}
	return set
}

// SliceUnion Union 两个 slice
func SliceUnion[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	union := aSet.Union(bSet)
	return union.ToSlice()
}

// SliceIntersection 两个切片的交集。提供的切片不需要是唯一的。不保证顺序。
func SliceIntersection[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	intersection := aSet.Intersection(bSet)
	return intersection.ToSlice()
}
