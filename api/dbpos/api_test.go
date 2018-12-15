package dbpos

import (
	"database/sql"
	"testing"
)

var (
	vide_id = ""
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate video_info")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

//按顺序执行
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("delete", testDeleteUser)
	t.Run("reget", testReget)
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("add", testAddNewVideo)
	t.Run("get", testGetVideo)
	t.Run("delete", testDeleteVideo)
	t.Run("reget", reGetVideo)

}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("idcpj", "123")
	if err != nil {
		t.Errorf("error of Adduser %v", err)
	}
}
func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("idcpj")
	if pwd != "123" || err != nil {
		t.Errorf("error of GetUSer %v", err)
	}
}
func testDeleteUser(t *testing.T) {
	err := DeleteUser("idcpj", "123")
	if err != nil {
		t.Errorf("error deleteUser %v", err)
	}
}

func testReget(t *testing.T) {
	pwd, err := GetUserCredential("idcpj")
	if pwd != "" || err != sql.ErrNoRows {
		t.Errorf("error of GetUSer %v", err)
	}
}

func testAddNewVideo(t *testing.T) {
	info, e := AddNewVideo(1, "test")
	if info == nil || e != nil {
		t.Errorf("info : %v, error:%v", info, e)
	}
	vide_id = info.Id
}

func testGetVideo(t *testing.T) {
	info, e := GetVideo(vide_id)
	if info == nil || e != nil {
		t.Errorf("info : %v , error :%v", info, e)
	}
}

func testDeleteVideo(t *testing.T) {
	e := DeleteVideo(vide_id)
	if e != nil {
		t.Errorf("error : %v", e)
	}
}
func reGetVideo(t *testing.T) {
	info, e := GetVideo(vide_id)
	if info != nil || e != sql.ErrNoRows {
		t.Errorf("info : %v , error :%v", info, e)
	}
}
