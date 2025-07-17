package main

// 单链表的增删改查实现

import "fmt"

// 泛型单链表节点
type SinglyLink[T comparable] struct {
	Next  *SinglyLink[T]
	Value T
}

//在链表头部插入节点
func (s *SinglyLink[T]) InsertAtHead(value T) *SinglyLink[T] {
	newNode := &SinglyLink[T]{Value: value, Next: s} //创建新节点并将其指向当前头部
	return newNode                                   //返回新头部
}

//在链表尾部插入节点
func (s *SinglyLink[T]) InsertAtTail(value T) *SinglyLink[T] {
	newNode := &SinglyLink[T]{Value: value}
	if s == nil {
		return newNode //如果链表为空，返回新节点作为头部
	}
	current := s
	for current.Next != nil {
		current = current.Next //遍历到链表尾部
	}
	current.Next = newNode //将新节点插入到尾部
	return s               //返回头部
}

//删除第一个匹配值的节点
func (s *SinglyLink[T]) Delete(value T) *SinglyLink[T] {
	if s == nil {
		return nil //如果链表为空，返回nil
	}
	current := s
	if current.Value == value { //如果头节点匹配，返回头节点的下一个节点
		return current.Next
	}
	for current.Next != nil {
		if current.Next.Value == value { //如果匹配，删除节点
			current.Next = current.Next.Next
			return s //返回头节点
		}
		current = current.Next //遍历到下一个节点
	}
	return s //如果没有匹配的节点，返回头节点
}

//修改第一个匹配值的节点
func (s *SinglyLink[T]) Update(value T, newValue T) *SinglyLink[T] {
	if s == nil {
		return nil //如果链表为空，返回nil
	}
	current := s
	if current.Value == value { //如果头节点匹配，更新头节点的值
		current.Value = newValue
		return s //返回头节点
	}
	for current.Next != nil {
		if current.Next.Value == value { //如果匹配，更新节点的值
			current.Next.Value = newValue
			return s //返回头节点
		}
		current = current.Next //遍历到下一个节点
	}
	return s //如果没有匹配的节点，返回头节点
}

// 查找节点是否存在
func (s *SinglyLink[T]) Exist(value T) bool {
	if s == nil {
		return false //如果链表为空，返回false
	}
	current := s
	for current != nil {
		if current.Value == value { //如果匹配，返回true
			return true
		}
		current = current.Next //遍历到下一个节点
	}
	return false //如果没有匹配的节点，返回false
}

// 打印链表
func (s *SinglyLink[T]) Print() {
	if s == nil {
		fmt.Println("Empty List")
		return
	}
	current := s
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
}

// 功能测试
func main() {
	var intList *SinglyLink[int] // 初始化空链表

	intList = intList.InsertAtHead(1) // 在头部插入1
	intList = intList.InsertAtTail(2) // 在尾部插入2
	intList = intList.InsertAtTail(3) // 在尾部插入3
	intList = intList.InsertAtTail(4) // 在尾部插入4

	intList.Print() // 打印链表
	fmt.Println()
	intList = intList.Update(4, 5) // 修改4为5
	intList.Print()
	fmt.Println()
	intList.Delete(5)                         // 删除5
	intList.Exist(3)                          // 检查3是否存在
	fmt.Println("Exist 3:", intList.Exist(3)) // 打印是否存在
	intList.Print()                           // 打印链表
}
