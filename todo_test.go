package todo

import (
	"testing"
)

func filledTodo() Todo {
	todo := NewTodo()

	item1 := NewItem()
	item2 := NewItem()

	item1.SetTitle("item1")
	item2.SetTitle("item2")

	todo.AddItems(item1, item2)
	todo.SetDone(0)
	todo.SetDone(0)
	todo.AddItems(item1, item2)

	return todo
}

func TestNewTodo(t *testing.T) {
	_ = NewTodo()
}

func TestAddItem(t *testing.T) {
	todo := NewTodo()

	item := NewItem()
	todo.AddItem(item)

	if len(todo.ActiveItems) != 1 {
		t.Fail()
	}
}

func TestAddItems(t *testing.T) {
	todo := NewTodo()

	item1 := NewItem()
	item2 := NewItem()

	todo.AddItems(item1, item2)

	if len(todo.ActiveItems) != 2 {
		t.Fail()
	}
}

func TestAddItemsAsSlice(t *testing.T) {
	todo := NewTodo()

	items := make(Items, 0)
	items = append(items, NewItem(), NewItem(), NewItem())
	todo.AddItems(items...)

	if len(todo.ActiveItems) != 3 {
		t.Fail()
	}
}

func TestGetItemByIndex(t *testing.T) {
	todo := filledTodo()
	item, err := todo.GetItem(1)

	if err != nil {
		t.Error(err)
	}

	if item.GetTitle() != "item2" {
		t.Fail()
	}
}

func TestGetItemByIndexAndExpectFail(t *testing.T) {
	todo := filledTodo()

	_, err := todo.GetItem(4)

	if err == nil {
		t.Error("No error reported")
	}
}

func TestSetDone(t *testing.T) {
	todo := filledTodo()

	err := todo.SetDone(1)

	if err != nil {
		t.Error(err)
	}
}

func TestSetDoneOutOfIndex(t *testing.T) {
	todo := filledTodo()

	err := todo.SetDone(100)

	if err == nil {
		t.Error("No error reported")
	}
}

func TestSetUndone(t *testing.T) {
	todo := filledTodo()
	err := todo.SetUndone(1)

	if err != nil {
		t.Error(err)
	}
}

func TestSetUndoneOutOfIndex(t *testing.T) {
	todo := filledTodo()

	err := todo.SetUndone(100)

	if err == nil {
		t.Error("No error reported")
	}
}

func TestGetItems(t *testing.T) {
	todo := filledTodo()
	items := todo.GetItems()

	if items == nil {
		t.Error("'items' is nil")
	}
}

func TestCount(t *testing.T) {
	todo := filledTodo()

	if todo.Count() != 2 {
		t.Error("Length is wrong")
	}
}

func TestRemove(t *testing.T) {
	todo := filledTodo()
	err := todo.Remove(0)

	if err != nil {
		t.Error(err)
	}
}
