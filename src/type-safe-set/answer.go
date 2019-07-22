package type_safe_set

import (
	"fmt"
	"reflect"
)

// Answer 是利用array或slice对Set的实现。
type Answer struct {
	elements []interface{}
	t        reflect.Type
}

// Size 返回TypeSafeSet中所包含的元素个数。
func (set *Answer) Size() int {
	return len(set.elements)
}

// IsEmpty 返回TypeSafeSet是否为空。
func (set *Answer) IsEmpty() bool {
	return len(set.elements) == 0
}

// 如果不存在返回 -1。
func (set *Answer) contains(e interface{}) (int, error) {
	if e == nil {
		return -1, fmt.Errorf("nil element")
	}
	if reflect.TypeOf(e) != set.t {
		return -1, fmt.Errorf("type error")
	}

	for i, each := range set.elements {
		if each == e {
			return i, nil
		}
	}
	return -1, nil
}

// Contains 判断TypeSafeSet中是否包含指定的元素e。
// 如果e的类型与TypeSafeSet元素的类型不一致返回error="type error"。
// 如果e为nil返回error="nil element"
func (set *Answer) Contains(e interface{}) (bool, error) {
	i, err := set.contains(e)
	if err != nil {
		return false, err
	}
	if i == -1 {
		return false, nil
	}
	return true, nil
}

// Add 添加一个原来没有的元素到TypeSafeSet中。
// 如果成功返回true。
// 如果e的类型与TypeSafeSet元素的类型不一致返回error="type error"。
// 如果e为nil返回error="nil element"。
func (set *Answer) Add(e interface{}) (bool, error) {
	if e == nil {
		return false, fmt.Errorf("nil element")
	}
	if set.IsEmpty() {
		set.t = reflect.TypeOf(e)
	}
	contains, err := set.Contains(e)
	if err != nil {
		return false, err
	}
	if contains {
		return false, nil
	}
	set.elements = append(set.elements, e)
	return true, nil
}

// Remove 删除TypeSafeSet中一个指定的元素.
// 如果元素存在返回true。
// 如果e的类型与TypeSafeSet元素的类型不一致返回error="type error"。
// 如果e为nil返回error="nil element"。
func (set *Answer) Remove(e interface{}) (bool, error) {
	i, err := set.contains(e)
	if err != nil {
		return false, err
	}
	if i == -1 {
		return false, nil
	}

	set.elements = append(set.elements[:i], set.elements[i+1:]...)
	return true, nil
}

// ToSlice 返回一个新的，包含有TypeSafeSet所有元素的slice
func (set *Answer) ToSlice() []interface{} {
	ret := make([]interface{}, len(set.elements))
	for i := range set.elements {
		ret[i] = set.elements[i]
	}
	return ret
}

// Equals判断该TypeSafeSet是否与另外一个TypeSafeSet相等。
// 两个TypeSafeSet相等意味着它们元素类型相同，个数相同，且彼此所包含的元素一样。
func (set *Answer) Equals(that TypeSafeSet) bool {
	if that == nil {
		return false
	}
	if set.Size() != that.Size() {
		return false
	}

	i := set.Iterator()
	for i.HasNext() {
		e, _ := i.Next()
		if contains, _ := that.Contains(e); !contains {
			return false
		}
	}
	return true
}

type arraySetIterator struct {
	Index int
	Size int
	Elements []interface{}
}

func (i *arraySetIterator) HasNext() bool {
	return i.Index != i.Size
}

func (i *arraySetIterator) Next() (interface{}, error) {
	if i.Index == i.Size {
		return nil, fmt.Errorf("out of range")
	}
	ret := i.Elements[i.Index]
	i.Index += 1
	return ret, nil

}

// Iterator 返回用于遍历TypeSafeSet元素的遍历器。
func (set *Answer) Iterator() Iterator {

	return &arraySetIterator{
		Index:    0,
		Size:     len(set.elements),
		Elements: set.ToSlice(),
	}
}
