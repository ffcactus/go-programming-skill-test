package type_safe_set

import (
	"fmt"
	"testing"
)

type BasicInterface interface {
	Print() string
}

type SuperInterface interface {
	BasicInterface
	PowerOff()
}

type BasicObject struct {
	Name string
}

func (BasicObject) Print() string {
	return "BasicObject"
}

type SuperObject struct {
	BasicObject
	Name string
}

func (SuperObject) Print() string {
	return "SuperObject"
}

func (SuperObject) PowerOff() {
}

func sameError(expected, actual error) bool {
	if expected == nil && actual == nil {
		return true
	}
	if expected != nil && actual == nil {
		return false
	}
	if expected == nil && actual != nil {
		return false
	}
	if expected.Error() != actual.Error() {
		return false
	}
	return true
}

// TestAnswer_Add 测试Add方法。
func TestAnswer_Add(t *testing.T) {
	samePointer := &BasicObject{}
	cases := []struct{
		description string
		set Answer
		firstElement interface{}
		secondElement interface{}
		success bool
		err error
	} {{
		description:   "add same type struct",
		set:           Answer{},
		firstElement:  BasicObject{Name: "first"},
		secondElement: BasicObject{Name: "second"},
		success:       true,
		err:           nil,
	}, {
		description:   "add different type struct",
		set:           Answer{},
		firstElement:  BasicObject{Name: "first"},
		secondElement: SuperObject{Name: "second"},
		success:       false,
		err:           fmt.Errorf("type error"),
	}, {
		description: "add same underlying type",
		set: Answer{},
		firstElement: &BasicObject{Name: "first"},
		secondElement: &BasicObject{Name: "second"},
		success: true,
		err: nil,
	}, {
		description: "add different underlying type",
		set: Answer{},
		firstElement: &BasicObject{},
		secondElement: &SuperObject{},
		success: false,
		err: fmt.Errorf("type error"),
	}, {
		description: "add duplicated element",
		set: Answer{},
		firstElement: 1,
		secondElement: 1,
		success: false,
		err: nil,
	}, {
		description: "add duplicated pointer",
		set: Answer{},
		firstElement: samePointer,
		secondElement: samePointer,
		success: false,
		err: nil,
	}, {
		description: "add nil",
		set: Answer{},
		firstElement: 1,
		secondElement: nil,
		success: false,
		err: fmt.Errorf("nil element"),
	}}

	for _, c := range cases {
		c.set.Add(c.firstElement)
		success, err := c.set.Add(c.secondElement)
		if (success != c.success) || (!sameError(err, c.err)) {
			t.Errorf("%s failed, expected=(%v, %v), actual=(%v, %v)",
				c.description,
				c.success, c.err,
				success, err)
		}
	}
}

// TestAnswer_Remove 测试Remove方法。
func TestAnswer_Remove(t *testing.T) {
	samePointer := &BasicObject{}
	cases := []struct{
		description string
		set Answer
		added interface{}
		toRemove interface{}
		success bool
		err error
	} {{
		description:   "remove exist element same type struct",
		set:           Answer{},
		added:  BasicObject{Name: "first"},
		toRemove: BasicObject{Name: "first"},
		success:       true,
		err:           nil,
	}, {
		description:   "remove different type struct",
		set:           Answer{},
		added:  BasicObject{Name: "first"},
		toRemove: SuperObject{Name: "first"},
		success:       false,
		err:           fmt.Errorf("type error"),
	}, {
		description: "remove same pointer",
		set: Answer{},
		added: samePointer,
		toRemove: samePointer,
		success: true,
		err: nil,
	}, {
		description: "remove non-exist element",
		set: Answer{},
		added: 1,
		toRemove: 2,
		success: false,
		err: nil,
	}, {
		description: "remove nil",
		set: Answer{},
		added: 1,
		toRemove: nil,
		success: false,
		err: fmt.Errorf("nil element"),
	}}

	for _, c := range cases {
		c.set.Add(c.added)
		success, err := c.set.Remove(c.toRemove)
		if (success != c.success) || (!sameError(err, c.err)) {
			t.Errorf("%s failed, expected=(%v, %v), actual=(%v, %v)",
				c.description,
				c.success, c.err,
				success, err)
		}
	}
}

func TestAnswer_Contains(t *testing.T) {
	samePointer := &BasicObject{}
	cases := []struct{
		description string
		set Answer
		added interface{}
		contained interface{}
		success bool
		err error
	} {{
		description:   "contains same type",
		set:           Answer{},
		added:  BasicObject{Name: "first"},
		contained: BasicObject{Name: "first"},
		success:       true,
		err:           nil,
	}, {
		description:   "contains different type struct",
		set:           Answer{},
		added:  BasicObject{Name: "first"},
		contained: SuperObject{Name: "first"},
		success:       false,
		err:           fmt.Errorf("type error"),
	}, {
		description: "contains same pointer",
		set: Answer{},
		added: samePointer,
		contained: samePointer,
		success: true,
		err: nil,
	}, {
		description: "contains different pointer",
		set: Answer{},
		added: &BasicObject{Name: "first"},
		contained: &BasicObject{Name: "first"},
		success: false,
		err: nil,
	}, {
		description: "contains nil",
		set: Answer{},
		added: 1,
		contained: nil,
		success: false,
		err: fmt.Errorf("nil element"),
	}}

	for _, c := range cases {
		c.set.Add(c.added)
		success, err := c.set.Contains(c.contained)
		if (success != c.success) || (!sameError(err, c.err)) {
			t.Errorf("%s failed, expected=(%v, %v), actual=(%v, %v)",
				c.description,
				c.success, c.err,
				success, err)
		}
	}
}

// TestAnswer_OperateWithFirstNil 测试第一次nil操作。
func TestAnswer_AddOrRemove_Nil(t *testing.T) {
	set1 := Answer{}
	set2 := Answer{}
	set3 := Answer{}
	added, err := set1.Add(nil)
	if added || err == nil {
		t.Errorf("Nil is not allowed. expect = (false, %v), actual = (%v, %v)",
			"nil element",
			added, err)
	}
	added, err = set2.Remove(nil)
	if added || err == nil {
		t.Errorf("Nil is not allowed. expect = (false, %v), actual = (%v, %v)",
			"nil element",
			added, err)
	}
	added, err = set3.Contains(nil)
	if added || err == nil {
		t.Errorf("Nil is not allowed. expect = (false, %v), actual = (%v, %v)",
			"nil element",
			added, err)
	}
}

// TestAnswer_Size 测试size方法。
func TestAnswer_Size(t *testing.T) {
	set := Answer{}
	if set.Size() != 0 {
		t.Errorf("Should be 0 size after initialized.")
	}
	set.Add(1)
	set.Add(2)
	if set.Size() != 2 {
		t.Errorf("Size doesn't increase normally.")
	}
	set.Remove(1)
	set.Remove(2)
	if set.Size() != 0 {
		t.Errorf("Size doesn't decrease normally.")
	}
}

// TestAnswer_IsEmpty 测试IsEmpty方法。
func TestAnswer_IsEmpty(t *testing.T) {
	set := Answer{}
	if ! set.IsEmpty() {
		t.Errorf("Should be empty after initialized.")
	}
	set.Add(1)
	if set.IsEmpty() {
		t.Errorf("IsEmpty doesn't work.")
	}
}

// TestAnswer_ToSlice 测试ToSlice方法。
func TestAnswer_ToSlice(t *testing.T) {
	set := Answer{}
	s1 := set.ToSlice()
	if s1 == nil {
		t.Errorf("ToSlice should not return nil")
	}
	if len(s1) != 0 {
		t.Errorf("ToSlice should not return zero sized slice if the set is empty.")
	}
	set.Add(1)
	set.Add(2)
	s2 := set.ToSlice()
	if len(s2) != 2 {
		t.Errorf("Size not match")
	}
}

// TestAnswer_Equals 测试Equals方法。
func TestAnswer_Equals(t *testing.T) {
	set1 := Answer{}
	if set1.Equals(nil) {
		t.Errorf("Equals to nil.")
	}
	set1.Add(1)
	set1.Add(2)

	set2 := Answer{}
	set2.Add(1)
	if set1.Equals(&set2) {
		t.Error("Equals doesn't work properly.")
	}
	set2.Add(2)
	if ! set1.Equals(&set2) {
		t.Error("Equals doesn't work properly.")
	}
	set2.Add(3)
	if set1.Equals(&set2) {
		t.Error("Equals doesn't work properly.")
	}
}

// TestAnswer_Iterator 测试Iterator方法。
func TestAnswer_Iterator(t *testing.T) {
	set := Answer{}

	i1 := set.Iterator()
	if i1.HasNext() {
		t.Errorf("Empty set should not have next value.")
	}
	_, err := i1.Next();
	if err == nil || err.Error() != "out of range"{
		t.Errorf("Next() should return out of range error when no more element, but got %v", err)
	}

	set.Add(1)
	set.Add(2)

	founded := 0
	i2 := set.Iterator()
	for i2.HasNext() {
		e, _ := i2.Next()
		if contained, _ := set.Contains(e); !contained {
			t.Error("Next() return elements that is not contained.")
		}
		founded += 1
	}
	if founded != set.Size() {
		t.Error("The elements returned by Iterator doesn't match Size().")
	}
}