package todo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Controller struct {
	repository repository
	todo       Todo
}

func NewController(path string) Controller {
	c := Controller{}
	c.repository = NewRepository(path)
	c.todo = c.repository.Load()
	return c
}

func (c *Controller) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()

		c.ShowItems()
		fmt.Println(" - - - - - - - - - - - - - ")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  add a\t\tAdds item(s) to the todo-list")
		fmt.Println("  remove r\tRemoves item(s) from the todo-list")
		fmt.Println("  clear c\tRemoves all items from the todo-list")
		fmt.Println("  exit e\tExits todo")
		fmt.Println("")
		fmt.Println(" - - - - - - - - - - - - - ")
		fmt.Println("")
		fmt.Printf("Enter command: ")
		text, _ := reader.ReadString('\n')

		text = strings.Trim(text, " \n")
		cmd := strings.Split(text, " ")

		if cmd[0] == "exit" || cmd[0] == "e" {
			break
		}
		if cmd[0] == "add" || cmd[0] == "a" {
			c.AddItems()
		}
		if cmd[0] == "remove" || cmd[0] == "r" {
			index, err := strconv.Atoi(cmd[1])
			if err == nil {
				c.RemoveItem(index - 1)
			}
		}
		if cmd[0] == "clear" || cmd[0] == "c" {
			c.Clear()
		}
	}
}

func (c *Controller) Clear() {
	c.todo.Clear()
	c.repository.Save(c.todo)
}

func (c *Controller) ShowItems() {
	fmt.Println("Todo")
	fmt.Println(" - - - - - - - - - - - - - ")
	for _, item := range c.todo.GetItems() {
		fmt.Println(" • " + item.Title)
	}
}

func (c *Controller) AddItems() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		if text == string('\n') {
			break
		}

		text = strings.Trim(text, " \n")

		item := NewItem()
		item.Title = text
		c.todo.AddItem(item)
		err := c.repository.Save(c.todo)

		if err != nil {
			panic(err)
		}
	}
}

func (c *Controller) RemoveItem(index int) {
	_ = c.todo.Remove(index)
	_ = c.repository.Save(c.todo)
}

func (c *Controller) AddItem(text string) {

}
