package task2

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var root *Node

	root = root.Insert(10)
	if root == nil || root.key[0] != 10 {
		t.Errorf("Корневой узел должен быть 10, но получено %v", root.key[0])
	}

	root = root.Insert(20)
	if root.size != 2 || root.key[1] != 20 {
		t.Errorf("Второй узел должен быть 20, но получено %v", root.key[1])
	}

	root = root.Insert(5)
	if root.size != 1 || root.first.key[0] != 5 || root.second.key[0] != 20 {
		t.Errorf("Узел должен быт разделена на 5 и 20, но получено %v и %v", root.first.key[0], root.second.key[0])
	}
}

func TestSearch(t *testing.T) {
	var root *Node

	root = root.Insert(10)
	root = root.Insert(20)
	root = root.Insert(5)
	root = root.Insert(6)
	root = root.Insert(15)

	result := root.Search(10)
	if result == nil || result.key[0] != 10 {
		t.Errorf("Ожидается 10, но получено %v", result.key[0])
	}

	result = root.Search(25)
	if result != nil {
		t.Errorf("Узел 25 не должн быть найден!")
	}
}

func TestRemove(t *testing.T) {
	var root *Node

	root = root.Insert(10)
	root = root.Insert(20)
	root = root.Insert(5)
	root = root.Insert(6)
	root = root.Insert(15)

	root = root.Remove(6)
	if root.Search(6) != nil {
		t.Errorf("Узел 6 удалён и не должен быть найден!")
	}

	root = root.Remove(10)
	if root.Search(10) != nil {
		t.Errorf("Узел 10 был удалён и не должен быть найден!")
	}
}
