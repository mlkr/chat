package data

import (
	"testing"
)

// 删除所有的帖子
func ThreadDeleteAll() (err error) {
	statement := "delete from threads"
	_, err = Db.Exec(statement)
	if err != nil {
		return
	}
	return
}

func Test_CreateThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "创建用户失败")
	}
	conv, err := users[0].CreateThread("第一次发帖")
	if err != nil {
		t.Error(err, "不能创建帖子")
	}
	if conv.UserId != users[0].Id {
		t.Error("用户没有与帖子关联")
	}
}
