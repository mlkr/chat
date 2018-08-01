package data

import (
	"fmt"
	"testing"
)

var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func setup() {
	// 删除所有帖子
	ThreadDeleteAll()

	// 删除所有的 session
	SessionDeleteAll()

	// 删除所有的用户
	UserDeleteAll()
}

func TestCreateUUID(t *testing.T) {
	fmt.Println(createUUID())
}
