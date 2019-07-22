package type_safe_set

// Iterator定义了便利集合的方法。
type Iterator interface {
	// HasNext 返回在当前位置之后是否还有下一个元素。
	HasNext() bool
	// Next 返回当前位置之后的下一个元素。如果没有返回error="out of range"
	Next() (interface{}, error)
}

// TypeSafeSet定义了一种集合，在这个集合中不允许有相同的元素。
type TypeSafeSet interface {
	// Size 返回TypeSafeSet中所包含的元素个数。
	Size() int

	// IsEmpty 返回TypeSafeSet是否为空。
	IsEmpty() bool

	// Contains 判断TypeSafeSet中是否包含指定的元素e。
	// 如果e的类型与TypeSafeSet元素的类型不一致返回error="type error"。
	// 如果e为nil返回error="nil element"
	Contains(e interface{}) (bool, error)

	// Iterator 返回用于遍历TypeSafeSet元素的遍历器。
	Iterator() Iterator

	// ToSlice 返回一个新的，包含有TypeSafeSet所有元素的slice。
	ToSlice() []interface{}

	// Add 添加一个原来没有的元素到TypeSafeSet中。
	// 如果成功返回true。
	// 如果e的类型与TypeSafeSet元素的类型不一致返回error="type error"。
	// 如果e为nil返回error="nil element"。
	Add(e interface{}) (bool, error)

	// Remove 删除TypeSafeSet中一个指定的元素.
	// 如果元素存在返回true。
	// 如果e的类型与TypeSafeSet元素的类型不一致返回error="type error"。
	// 如果e为nil返回error="nil element"。
	Remove(e interface{}) (bool, error)

	// Equals判断该TypeSafeSet是否与另外一个TypeSafeSet相等。
	// 两个TypeSafeSet相等意味着它们元素类型相同，个数相同，且彼此所包含的元素一样。
	// 如果另一个为nil则不相等。
	Equals(that TypeSafeSet) bool
}