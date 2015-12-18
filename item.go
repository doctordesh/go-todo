package todo

import (
	"errors"
	"time"
)

type Item struct {
	Title    string     `json:"title"`
	Done     bool       `json:"done"`
	DateDone *time.Time `json:"date_done"`
}

func NewItem() Item {
	item := Item{}
	item.Done = false
	return item
}

func (i *Item) SetTitle(t string) {
	i.Title = t
}

func (i *Item) GetTitle() string {
	return i.Title
}

func (i *Item) Complete() {
	now := time.Now()
	i.Done = true
	i.DateDone = &now
}

func (i *Item) Open() {
	i.Done = false
	i.DateDone = nil
}

func (i *Item) IsDone() bool {
	return i.Done
}

func (i *Item) CompletedAt() *time.Time {
	return i.DateDone
}

type Items []Item

func (i *Items) RemoveAt(index int) (Item, error) {
	if len(*i) < index || index < 0 {
		return Item{}, errors.New("Index out of range")
	}

	item := (*i)[index]
	tmp := append((*i)[:index], (*i)[index+1:]...)
	*i = tmp

	return item, nil
}
