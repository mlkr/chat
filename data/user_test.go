package data

import (
	"database/sql"
	"testing"
)

func Test_UserCreate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	if users[0].Id == 0 {
		t.Errorf("创建用户 id 错误")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "创建用户错误, 没有找到该 email 的用户")
	}
	if users[0].Email != u.Email {
		t.Errorf("创建用户错误 email 不一样")
	}
}

func Test_UserDelete(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	if err := users[0].Delete(); err != nil {
		t.Error(err, "- 不能删除用户")
	}
	_, err := UserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- 用户没有删除")
	}
}

func Test_UserUpdate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	users[0].Name = "Random User"
	if err := users[0].Update(); err != nil {
		t.Error(err, "- 不能更新用户")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "- 不能获取更新后的用户")
	}
	if u.Name != "Random User" {
		t.Error(err, "- 用户没有更新")
	}
}

func Test_UserByUUID(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	u, err := UserByUUID(users[0].Uuid)
	if err != nil {
		t.Error(err, "不能通过 UUID 获取用户")
	}
	if users[0].Email != u.Email {
		t.Errorf("创建的用户的 email 和原来的不一致")
	}
}

func Test_Users(t *testing.T) {
	setup()
	for _, user := range users {
		if err := user.Create(); err != nil {
			t.Error(err, "不能创建用户")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "不能获取所有用户")
	}
	if len(u) != 2 {
		t.Error(err, "获取的用户数量不对")
	}
	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "获取的用户 email 和原来的不一致")
	}
}

func Test_CreateSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "不能创建 session")
	}
	if session.UserId != users[0].Id {
		t.Error("创建的 session 没有和用户联系一起")
	}
}

func Test_GetSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "不能创建 session")
	}

	s, err := users[0].Session()
	if err != nil {
		t.Error(err, "获取 session 失败")
	}
	if s.Id == 0 {
		t.Error("没有获取到 session")
	}
	if s.Id != session.Id {
		t.Error("获取的 session 与创建的不一致")
	}
}

func Test_checkValidSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "不能创建 session")
	}

	uuid := session.Uuid

	s := Session{Uuid: uuid}
	valid, err := s.Check()
	if err != nil {
		t.Error(err, "检查 session 失败")
	}
	if valid != true {
		t.Error(err, "session 是非法的")
	}

}

func Test_checkInvalidSession(t *testing.T) {
	setup()
	s := Session{Uuid: "123"}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "获取 session 的 sql 语句有问题")
	}
	if valid == true {
		t.Error(err, "session 的非法检查有问题")
	}

}

func Test_DeleteSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "不能创建用户")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "不能创建 session")
	}

	err = session.DeleteByUUID()
	if err != nil {
		t.Error(err, "不能删除 session")
	}
	s := Session{Uuid: session.Uuid}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "session 被删除后还能检查出错误")
	}
	if valid == true {
		t.Error(err, "session 没有被删除")
	}
}
