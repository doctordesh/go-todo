package todo

import (
	"fmt"
	"testing"
)

func TestNewItem(t *testing.T) {
	_ = NewItem()
}

func TestSetTitle(t *testing.T) {
	item := NewItem()
	title := "This is the title"
	item.SetTitle(title)
	if item.GetTitle() != title {
		t.Fail()
	}
}

func TestComplete(t *testing.T) {
	item := NewItem()
	item.SetTitle("This is a task")

	if item.IsDone() == true {
		t.Fail()
	}

	item.Complete()

	if item.IsDone() == false {
		t.Fail()
	}

	if item.CompletedAt() == nil {
		t.Fail()
	}
}

func TestOpen(t *testing.T) {
	item := NewItem()
	item.Open()
}

func TestItems(t *testing.T) {
	fmt.Println()
	items := Items{NewItem(), NewItem(), NewItem()}
	length := len(items)
	_, _ = items.RemoveAt(1)
	if length-1 != len(items) {
		t.Logf("Length before: %d, length after: %d\n", length, len(items))
		t.Error("Length of Items not changed after Items.RemoveAt(index)")
	}
}

func TestItemsOutOfRande(t *testing.T) {
	items := Items{NewItem(), NewItem(), NewItem()}
	_, err := items.RemoveAt(-1)
	if err == nil {
		t.Error("Should return error on remove out of range")
	}
}
