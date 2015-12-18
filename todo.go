package todo

import (
	"errors"
	"fmt"
)

type Todo struct {
	ActiveItems Items `json:"active_items"`
	DoneItems   Items `json:"done_items"`
}

func NewTodo() Todo {
	return Todo{}
}

func (t *Todo) GetItems() Items {
	return t.ActiveItems
}

func (t *Todo) AddItems(items ...Item) {
	t.ActiveItems = append(t.ActiveItems, items...)
}

func (t *Todo) AddItem(item Item) {
	t.ActiveItems = append(t.ActiveItems, item)
}

func (t *Todo) Remove(index int) error {
	_, err := t.ActiveItems.RemoveAt(index)
	return err
}

func (t *Todo) Count() int {
	return len(t.ActiveItems)
}

func (t *Todo) GetItem(index int) (Item, error) {
	if index >= len(t.ActiveItems) {
		return Item{}, indexOutOfBoundsError(index)
	}
	return t.ActiveItems[index], nil
}

func (t *Todo) SetDone(index int) error {
	item, err := t.ActiveItems.RemoveAt(index)
	if err != nil {
		return err
	}
	item.Complete()
	t.DoneItems = append(t.DoneItems, item)
	return nil
}

func (t *Todo) SetUndone(index int) error {
	item, err := t.DoneItems.RemoveAt(index)
	if err != nil {
		return err
	}
	item.Open()
	t.ActiveItems = append(t.ActiveItems, item)
	return nil
}

func (t *Todo) Clear() {
	t.ActiveItems = Items{}
	t.DoneItems = Items{}
}

func indexOutOfBoundsError(index int) error {
	return errors.New(fmt.Sprintf("Index (%d) out of bounds", index))
}
