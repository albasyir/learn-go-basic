package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Less(a int, b int) bool {
	return users[a].Age < users[b].Age
}

func (users Users) Swap(a int, b int) {
	users[a], users[b] = users[b], users[a]
}

func main() {

	users := Users{
		{"BBB", 20},
		{"AAA", 10},
		{"DDD", 40},
		{"CCC", 30},
	}

	fmt.Println(users)
	sort.Sort(users)
	fmt.Println(users)
}
