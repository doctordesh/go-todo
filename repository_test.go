package todo

import (
	"os/exec"
	"testing"
)

const PATH = "/Users/emil/todo.json"

func setup() {
	exec.Command("rm", PATH)
}

func loadRepository(t *testing.T) (repository, Todo) {
	repository := NewRepository(PATH)
	todo := repository.Load()
	return repository, todo
}

func TestEmptyLoad(t *testing.T) {
	setup()
	_, _ = loadRepository(t)
}

func TestSave(t *testing.T) {
	setup()
	repository, todo := loadRepository(t)
	item := NewItem()
	item.SetTitle("Lorem ipsum")
	todo.AddItem(item)
	err := repository.Save(todo)
	if err != nil {
		t.Error(err)
	}
}
