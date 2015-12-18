package todo

import (
	"encoding/json"
	"io/ioutil"
)

type repository struct {
	Path string
}

func NewRepository(path string) repository {
	return repository{Path: path}
}

func (c *repository) Load() Todo {
	data, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return NewTodo()
	}
	todo := NewTodo()
	err = json.Unmarshal(data, &todo)
	if err != nil {
		return NewTodo()
	}
	return todo
}

func (c *repository) Save(todo Todo) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.Path, data, 0755)
}
