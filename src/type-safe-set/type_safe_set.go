// 一个Set实现的测试。
// Set是指不包含相同元素的集合。但是你的Set的实现需要额外满足一下一些要求：
// 当先该Set添加第一个元素时，该Set所能包含的元素类型就确定了。
// 具体来就是元素的reflect.TypeOf()必须相同。
// 不能添加nil。
// interface.go中包含了Set的接口定义，type_safe_set_test.go包含了测试用例。你的实现需要写在type_safe_set.go中。
package type_safe_set
